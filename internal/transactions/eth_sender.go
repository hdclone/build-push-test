package transactions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"broadcaster/internal/bridge"
	"broadcaster/internal/config"
	"broadcaster/internal/logging"
	"broadcaster/internal/model"
	"broadcaster/internal/modules"
	repository "broadcaster/internal/repository/transactions"
)

func EthSend(
	ctx context.Context,
	transaction *model.Transaction,
	chainConfig *config.ChainConfig,
	transactOpts *bind.TransactOpts,
	logger *zap.Logger,
) (retryErr error) {
	advisorResult, err := requestGasAdvisor(ctx, transaction, logger)
	if err != nil {
		return err
	}

	clients := modules.Pool().Clients(chainConfig)

	// prepare options of transaction
	transactOpts.Value = big.NewInt(0)
	transactOpts.GasPrice = big.NewInt(advisorResult.GasPrice)
	transactOpts.GasPrice.SetUint64(transactOpts.GasPrice.Uint64() + (transactOpts.GasPrice.Uint64()/100)*uint64(CtxGetGasIncrementByPercents(ctx)))
	if chainConfig.GasLimit > 0 {
		logger.Warn(fmt.Sprintf("use gas limit %d from config", chainConfig.GasLimit))
		transactOpts.GasLimit = chainConfig.GasLimit
	} else {
		var gas uint64
		retryErr = clients.Try(ctx, func(ctx context.Context, client *ethclient.Client) (estimateGasErr error) {
			gas, estimateGasErr = EstimateGas(ctx, chainConfig.Key.Address, transaction.BridgeAddress, "receiveRequestV2Signed", client, transaction.CallData, transaction.ReceiveSide, transaction.Signature)
			if estimateGasErr != nil {
				return errors.WithMessage(estimateGasErr, "EstimateGas")
			}
			return
		})
		if retryErr != nil {
			return retryErr
		}
		if gas < 1000000 {
			transactOpts.GasLimit = 1000000
			logger.Warn(fmt.Sprintf("estimated gas %d was changed to %d", gas, transactOpts.GasLimit))
		} else {
			transactOpts.GasLimit = gas * 120 / 100
		}
	}

	// try to get nonce
	if pendingNonceErr := clients.Try(ctx, func(ctx context.Context, client *ethclient.Client) error {
		pendingNonce, pendingNonceAtErr := client.PendingNonceAt(context.TODO(), chainConfig.Key.Address)
		if pendingNonceAtErr != nil {
			return errors.WithMessage(pendingNonceAtErr, "PendingNonceAt")
		} else {
			transactOpts.Nonce = big.NewInt(int64(pendingNonce))
		}
		return nil
	}); pendingNonceErr != nil {
		return pendingNonceErr
	}

	if advisorResult.GasPriceE1559 {
		transactOpts.GasFeeCap = transactOpts.GasPrice
		transactOpts.GasTipCap = big.NewInt(advisorResult.GasPricePriority)
		transactOpts.GasPrice = nil
	}

	// set transaction status to sending
	sendingStatusErr := repository.SetSendingStatus(ctx, transaction, transactOpts.Nonce.Uint64(), transactOpts.GasLimit, transactOpts.GasPrice)

	if sendingStatusErr != nil {
		return errors.WithMessage(sendingStatusErr, "can't set sending status for transaction")
	}

	logger.Info(fmt.Sprintf("send transaction (nonce: %v, gasLimit: %v, eip1559: %v, gasPrice: %v, gasFeeCap: %v, gasTipCap: %v)", transactOpts.Nonce.Uint64(), transactOpts.GasLimit, advisorResult.GasPriceE1559, transactOpts.GasPrice, transactOpts.GasFeeCap, transactOpts.GasTipCap))

	var trxSuccessfullySent *types.Transaction

	if retryErr = clients.Try(ctx, func(ctx context.Context, client *ethclient.Client) (pendingNonceAtErr error) {
		contract, err := bridge.NewBridge(transaction.BridgeAddress, client)
		if err != nil {
			return errors.WithMessage(err, "can't create bridge contract")
		}

		trx, err := contract.ReceiveRequestV2Signed(
			transactOpts,
			transaction.CallData,
			transaction.ReceiveSide,
			transaction.Signature,
		)
		if err != nil {
			return errors.WithMessage(err, "can't call ReceiveRequestV2Signed")
		}

		if trxSuccessfullySent == nil && trx != nil {
			trxSuccessfullySent = trx
		}

		if trxSuccessfullySent == nil && trx != nil {
			trxSuccessfullySent = trx
		}

		return nil
	}); retryErr != nil {
		return retryErr
	}

	if trxSuccessfullySent != nil {
		return repository.SetSentStatus(logging.CtxSet(ctx, logger.With(zap.String("tx_hash", trxSuccessfullySent.Hash().String()))), transaction, trxSuccessfullySent.Hash().Bytes())
	}
	return nil
}
