package broadcaster_client_schema

import (
	broadcaster_client_errors "broadcaster/client/errors"
	"net/http"

	eth_common "github.com/ethereum/go-ethereum/common"
)

type RequestBroadcast struct {
	CallData    []byte             `json:"call_data"`
	ReceiveSide eth_common.Address `json:"receive_side"`
	Signature   []byte             `json:"signature"`
}

func (request *RequestBroadcast) Bind(r *http.Request) error {
	if len(request.Signature) == 0 {
		return broadcaster_client_errors.BroadcastRequestSignatureEmpty
	}

	return nil
}
