package transactions

import (
	"broadcaster/internal/bridge"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func EstimateGas(ctx context.Context, from common.Address, to common.Address, method string,
	transactor bind.ContractTransactor, parameters ...interface{},
) (uint64, error) {
	contractABI, err := bridge.BridgeMetaData.GetAbi()
	if err != nil {
		return 0, err
	}

	input, err := contractABI.Pack(method, parameters...)
	if err != nil {
		return 0, err
	}

	msg := ethereum.CallMsg{
		From: from,
		To:   &to,
		Data: input,
	}

	gas, err := transactor.EstimateGas(ctx, msg)
	if err != nil {
		return 0, err
	}
	return gas, nil
}
