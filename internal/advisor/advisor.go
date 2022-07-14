package advisor

import (
	"broadcaster/internal/config"
	"broadcaster/internal/model"
	"broadcaster/internal/server"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"net/http"
)

//go:generate minimock -i GasPrice -o ./ -s _mock.go -g

type ResponseResult string

const (
	ResultAccept = ResponseResult("accept")
	ResultReject = ResponseResult("reject")
)

func (rr ResponseResult) Accepted() bool {
	return rr == ResultAccept
}

func (rr ResponseResult) Rejected() bool {
	return rr == ResultReject
}

type Response struct {
	Result           ResponseResult `json:"result"`
	GasPrice         int64          `json:"gas_price"`
	GasPriceE1559    bool           `json:"gas_price_e1559"`
	GasPricePriority int64          `json:"gas_price_priority"`
}

type RequestParams struct {
	CallData    string         `json:"call_data"`
	ChainIDFrom int64          `json:"chain_id_from"`
	ChainIDTo   int64          `json:"chain_id_to"`
	ReceiveSide common.Address `json:"receive_side"`
	TxHash      common.Hash    `json:"tx_hash"`
}

type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (re ResponseError) Error() string {
	if re.Details != nil {
		return fmt.Sprintf("%s (details: %v)", re.Message, re.Details)
	}
	return re.Message
}

type Interface interface {
	RequestGasPrice(ctx context.Context, trans *model.Transaction) (*Response, error)
}

type Advisor struct {
	url    string
	client *http.Client
}

type Mock struct {
	config *config.AdvisorMockConfig
}

func (m Mock) RequestGasPrice(_ context.Context, _ *model.Transaction) (*Response, error) {
	return &Response{
		Result:           ResponseResult(m.config.Result),
		GasPricePriority: m.config.GasPricePriority,
		GasPriceE1559:    m.config.GasPriceE1559,
		GasPrice:         m.config.GasPrice,
	}, nil
}

func NewAdvisorMock(mock *config.AdvisorMockConfig) Interface {
	return &Mock{config: mock}
}

func NewAdvisor(url string, client *http.Client) Interface {
	return &Advisor{url: url, client: client}
}

// RequestGasPrice send request to microservice for validate transaction and get gas price
func (a *Advisor) RequestGasPrice(ctx context.Context, trans *model.Transaction) (*Response, error) {
	var (
		requestId      = server.CtxGetRequestID(ctx)
		responseParams = &Response{}
	)

	requestParams := &RequestParams{
		TxHash:      common.BytesToHash(trans.HashSource),
		CallData:    "0x" + hex.EncodeToString(trans.CallData),
		ChainIDFrom: trans.ChainIDFrom,
		ChainIDTo:   trans.ChainID,
		ReceiveSide: trans.ReceiveSide,
	}

	requestBody := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(requestBody).Encode(requestParams); err != nil {
		return nil, errors.WithMessage(err, "can't encode request body")
	}

	request, err := http.NewRequest(http.MethodPost, a.url, requestBody)
	if err != nil {
		return nil, errors.WithMessage(err, "can't create request")
	}
	request.Header.Add("Content-Type", "application/json")
	if requestId != nil {
		request.Header.Add("X-Request-ID", requestId.String())
	}

	resp, err := a.client.Do(request)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to send request")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		var re ResponseError
		if errDecode := json.NewDecoder(resp.Body).Decode(&re); errDecode != nil {
			return nil, errors.WithMessage(errDecode, "can't decode error response")
		}
		return nil, re
	}

	if err = json.NewDecoder(resp.Body).Decode(responseParams); err != nil {
		err = errors.WithMessage(err, "failed to decode response")
		return nil, err
	}

	return responseParams, nil
}
