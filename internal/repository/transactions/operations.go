package transactions

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"broadcaster/internal/database"
	"broadcaster/internal/logging"
	"broadcaster/internal/model"
)

func Insert(ctx context.Context, transaction *model.Transaction) error {
	transaction.CreatedAt = model.PtrTime(time.Now())
	_, err := database.CtxGet(ctx).NewInsert().Model(transaction).Exec(ctx)
	return err
}

func Update(ctx context.Context, transaction *model.Transaction) error {
	transaction.UpdatedAt = model.PtrTime(time.Now())
	_, err := database.CtxGet(ctx).NewUpdate().Model(transaction).WherePK().Exec(ctx)
	return err
}

func Create(ctx context.Context,
	chainIDFrom int64,
	chainID int64,
	bridgeAddress common.Address,
	receiveSide common.Address,
	callData []byte,
	signature []byte,
	hashSource common.Hash,
	senderAddress common.Address,
	networkType model.NetworkType,
) (*model.Transaction, error) {
	transaction := &model.Transaction{
		ID:            model.NewID("trx"),
		ChainIDFrom:   chainIDFrom,
		ChainID:       chainID,
		BridgeAddress: bridgeAddress,
		ReceiveSide:   receiveSide,
		CallData:      callData,
		Signature:     signature,
		HashSource:    hashSource.Bytes(),
		SenderAddress: senderAddress,
		Status:        model.TransactionStatusPending,
		NetworkType:   &networkType,
	}

	// https://www.google.com/search?q=ecrecover+27+28
	if transaction.Signature[len(transaction.Signature)-1] <= 1 {
		transaction.Signature[len(transaction.Signature)-1] += 27
	}

	err := Insert(ctx, transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func SetPendingStatus(ctx context.Context, transaction *model.Transaction) error {
	logging.CtxGet(ctx).Info("set pending status")
	transaction.Status = model.TransactionStatusPending
	transaction.Nonce = nil
	return Update(ctx, transaction)
}

func SetSendingStatus(ctx context.Context, transaction *model.Transaction, nonce uint64, gasLimit uint64, gasPrice *big.Int) error {
	logging.CtxGet(ctx).Info("set sending status")
	transaction.Nonce = &nonce
	transaction.GasLimit = &gasLimit
	transaction.GasPrice = (*model.BigInt)(gasPrice)
	transaction.Status = model.TransactionStatusSending
	transaction.Error = nil
	return Update(ctx, transaction)
}

func SetSentStatus(ctx context.Context, transaction *model.Transaction, hash []byte) error {
	logging.CtxGet(ctx).Info("set sent status")
	transaction.Status = model.TransactionStatusSent
	transaction.Hash = hash
	return Update(ctx, transaction)
}

func SetFailedStatus(ctx context.Context, transaction *model.Transaction, err error) error {
	logging.CtxGet(ctx).Info("set failed status")
	transaction.Status = model.TransactionStatusFailed
	transaction.Error = model.PtrString(err.Error())
	return Update(ctx, transaction)
}
