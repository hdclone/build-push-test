package main

import (
	broadcaster_client_schema "broadcaster/client/schema"
	"broadcaster/cmd/test-tool/mock"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/docopt/docopt-go"
	eth_bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	eth_client "github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var version = "[manual build]"

var usage = `
Usage:
  test-tool-tool -h | --help
  test-tool-tool --private-key=<key> (
                deploy-mock-contract --network=<endpoint> |
                sign --call-data=<data> |
                request --call-data=<data> --endpoint=<endpoint> |
                siege --endpoint=<endpoint> --requests=<amount> --concurrency=<limit>
            )

Options:
  -h --help  Show this help.
`

// "wss://rinkeby.infura.io/ws/v3/b3572a473b864c489a470acf49d4b41b"
// "bfe415cb5bff47d5d9eb45ffc9d06bf757287dd2c7c44b81debc434942e1d25c"

// data hash 0x38a3202b1fe5f648045243851fe49ec10d3ddc41a334a481f9417eb92ec17a6e
// pub key 0x32F1d8C86ba575EaC312E5712e01d912B8234528
// signature gCh/co3VQnoIog2dgF6F7EB+Od1P2tA64fJKA3dL1XFWTWniUXo6yf7CAeVMaNnjaPKOBHnTl23y5EfUUmO7fgE=
// balance 97309409165175766
// trx id 0xa6213bdd836f0a35c2fd211b129e88eb9fb51b11a87ff8e1bbd25c9a1d23ba51
// contract address 0x2289d3279494198254A02Fa3DDB721DB454e8a74

type Opts struct {
	ValueNetwork              string `docopt:"--network"`
	ValuePrivateKey           string `docopt:"--private-key"`
	ValueCallData             string `docopt:"--call-data"`
	ValueAuth                 string `docopt:"--auth"`
	ValueEndpoint             string `docopt:"--endpoint"`
	ValueRequests             int    `docopt:"--requests"`
	ValueConcurrency          int    `docopt:"--concurrency"`
	CommandDeployMockContract bool   `docopt:"deploy-mock-contract"`
	CommandSign               bool   `docopt:"sign"`
	CommandRequest            bool   `docopt:"request"`
	CommandSiege              bool   `docopt:"siege"`
}

func main() {
	logger, _ := zap.NewDevelopment()

	args, err := docopt.ParseArgs(usage, nil, "broadcaster "+version)
	if err != nil {
		logger.Fatal(err.Error())
	}

	var opts Opts
	err = args.Bind(&opts)
	if err != nil {
		logger.Fatal(err.Error())
	}

	privateKey, err := crypto.HexToECDSA(opts.ValuePrivateKey)
	if err != nil {
		logger.Fatal(err.Error())
	}

	publicKeyECDSA := *privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(publicKeyECDSA)

	switch {
	case opts.CommandSiege:
		var (
			wg       = sync.WaitGroup{}
			requests = make(chan *http.Request)
		)

		for worker := 0; worker <= opts.ValueConcurrency; worker++ {
			go func(worker int) {
				for request := range requests {
					request.Header.Add("content-type", "application/json")

					auth := strings.SplitN(opts.ValueAuth, ":", 2)
					if len(auth) == 2 {
						request.SetBasicAuth(auth[0], auth[1])
					}

					response, err := http.DefaultClient.Do(request)
					if err != nil {
						logger.Fatal(err.Error())
					} else {
						fmt.Println(worker, response.Status)
					}

					wg.Done()
				}
			}(worker)
		}

		for i := 0; i < opts.ValueRequests; i++ {
			var callData [64]byte
			_, err := rand.Read(callData[:])
			if err != nil {
				logger.Fatal(err.Error())
			}

			request, err := signRequest(callData[:], fromAddress, privateKey)
			if err != nil {
				logger.Fatal(err.Error())
			}

			data, err := json.Marshal(request)
			if err != nil {
				logger.Fatal(err.Error())
			}

			post, err := http.NewRequest(
				http.MethodPost,
				opts.ValueEndpoint,
				bytes.NewReader(data),
			)
			if err != nil {
				logger.Fatal(err.Error())
			}

			wg.Add(1)
			requests <- post
		}

		wg.Wait()
	case opts.CommandRequest:
		callData := []byte(opts.ValueCallData)

		request, err := signRequest(callData, fromAddress, privateKey)
		if err != nil {
			logger.Fatal(err.Error())
		}

		data, err := json.Marshal(request)
		if err != nil {
			logger.Fatal(err.Error())
		}

		post, err := http.NewRequest(
			http.MethodPost,
			opts.ValueEndpoint,
			bytes.NewReader(data),
		)
		if err != nil {
			logger.Fatal(err.Error())
		}

		post.Header.Add("content-type", "application/json")

		auth := strings.SplitN(opts.ValueAuth, ":", 2)
		if len(auth) == 2 {
			post.SetBasicAuth(auth[0], auth[1])
		}

		response, err := http.DefaultClient.Do(post)
		if err != nil {
			logger.Fatal(err.Error())
		}

		body, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("RequestID: %s\nStatus: %v\nBody: %v", response.Header.Get("X-Request-ID"), response.Status, string(body))

	case opts.CommandSign:
		callData := []byte(opts.ValueCallData)

		request, err := signRequest(callData, fromAddress, privateKey)
		if err != nil {
			logger.Fatal(err.Error())
		}

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(request); err != nil {
			logger.Fatal(err.Error())
		}

	case opts.CommandDeployMockContract:
		client, err := eth_client.Dial(opts.ValueNetwork)
		if err != nil {
			logger.Fatal(err.Error())
		}

		balance, err := client.BalanceAt(
			context.TODO(),
			crypto.PubkeyToAddress(publicKeyECDSA),
			nil,
		)
		if err != nil {
			logger.Fatal(err.Error())
		}

		fmt.Println("balance", balance)

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			logger.Fatal(err.Error())
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			logger.Fatal(err.Error())
		}

		auth := eth_bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = 3000000
		auth.GasPrice = gasPrice

		address, trx, _, err := mock.DeployMock(auth, client, fromAddress)
		if err != nil {
			logger.Fatal(err.Error())
		}

		fmt.Println("trx hash", trx.Hash().Hex())
		fmt.Println("contract address", address.Hex())
	}
}

func signRequest(
	callData []byte,
	fromAddress eth_common.Address,
	privateKey *ecdsa.PrivateKey,
) (*broadcaster_client_schema.RequestBroadcast, error) {
	data := append([]byte("receiveRequestV2"), callData...)
	data = append(data, fromAddress.Bytes()...)
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	return &broadcaster_client_schema.RequestBroadcast{
		CallData:    []byte(callData),
		ReceiveSide: fromAddress,
		Signature:   signature,
	}, nil
}
