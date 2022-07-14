package transactions

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"broadcaster/internal/advisor"
	"broadcaster/internal/config"
	"broadcaster/internal/logging"
	"broadcaster/internal/model"
	"broadcaster/internal/modules"
	repository "broadcaster/internal/repository/transactions"
)

const maxProcessingDelay = 95

type incrementContextKey struct{}

var ick = incrementContextKey{}

func CtxSetGasIncrementByPercents(ctx context.Context, incrementBy uint) context.Context {
	return context.WithValue(ctx, ick, &incrementBy)
}

func CtxGetGasIncrementByPercents(ctx context.Context) uint {
	if incrementBy := ctx.Value(ick); incrementBy != nil {
		return *(incrementBy).(*uint)
	}
	return 0
}

func Sender(ctx context.Context, transaction *model.Transaction, withRetry bool) (senderErr error) {
	var (
		chainConfig  *config.ChainConfig
		retryOptions []retry.Option
	)

	logger := modules.Logger().With(zap.String("tx_hash_source", "0x"+common.Bytes2Hex(transaction.HashSource)))
	ctx = logging.CtxSet(ctx, logger)
	ctx = CtxSetGasIncrementByPercents(ctx, 0)

	sourceTxDuration(ctx, transaction, logger)

	defer func() {
		if senderErr != nil {
			setStatusErr := repository.SetFailedStatus(ctx, transaction, senderErr)
			if setStatusErr != nil {
				logger.Error(setStatusErr.Error())
			}
		}
	}()

	// can be send transactions with failed or pending status only
	if !(transaction.Status == model.TransactionStatusFailed || transaction.Status == model.TransactionStatusPending) {
		logger.Warn(fmt.Sprintf("can't send transaction with '%s' status", transaction.Status))
		return
	}

	chainConfig, senderErr = config.CtcGet(ctx).Chains.Get(transaction.ChainID)
	if senderErr != nil {
		return
	}

	if withRetry {
		retryOptions = []retry.Option{
			retry.Attempts(chainConfig.Retry.Attempts),
			retry.Delay(chainConfig.Retry.Delay.Duration()),
			retry.DelayType(retry.FixedDelay),
			retry.OnRetry(func(attempt uint, err error) {
				// TODO: need to improve handling transaction underpriced error
				if strings.Contains(err.Error(), "transaction underpriced") { // transaction underpriced
					percents := CtxGetGasIncrementByPercents(ctx) + 10
					logger.With(zap.Uint("attempt", attempt+1)).Warn(fmt.Sprintf("increment gas by %v percent", percents))
					ctx = CtxSetGasIncrementByPercents(ctx, percents)
				}
				logger.With(zap.Uint("attempt", attempt+1)).Error(err.Error())
			}),
		}
	} else {
		retryOptions = []retry.Option{
			retry.Attempts(1),
		}
	}

	if chainConfig.Key.Key == nil {
		senderErr = fmt.Errorf("key for chain %v is not configured", chainConfig.ID)
		return
	}

	if transaction.NetworkType == nil {
		return fmt.Errorf("undefined chain type")
	}

	if strings.EqualFold(string(*transaction.NetworkType), string(model.ChainTypeEvm)) {
		var transactOpts *bind.TransactOpts
		transactOpts, senderErr = bind.NewKeyedTransactorWithChainID(chainConfig.Key.Key, big.NewInt(chainConfig.ID))
		if senderErr != nil {
			return
		}
		senderErr = retry.Do(func() error {
			return EthSend(ctx, transaction, chainConfig, transactOpts, logger)
		}, retryOptions...)
	} else if *transaction.NetworkType == model.ChainTypeCosmos {
		senderErr = retry.Do(func() error {
			return TerraSend(ctx, transaction, chainConfig, logger)
		}, retryOptions...)
	} else {
		senderErr = fmt.Errorf("wrong chain type, received: %s", *transaction.NetworkType)
	}

	return
}

func requestGasAdvisor(ctx context.Context, transaction *model.Transaction, logger *zap.Logger) (*advisor.Response, error) {
	advisorResult, err := modules.Advisor().RequestGasPrice(ctx, transaction)
	if err != nil {
		return nil, errors.WithMessage(err, "advisor")
	} else if !advisorResult.Result.Accepted() {
		return nil, errors.New("rejected transaction by advisor")
	}
	logger.Info(fmt.Sprintf("advisor result (Price: %v, E1559: %v, Priority: %v)", advisorResult.GasPrice, advisorResult.GasPriceE1559, advisorResult.GasPricePriority))
	return advisorResult, nil
}

//sourceTxDuration calculate duration from source tx block mined to receive tx from relayers
func sourceTxDuration(ctx context.Context, transaction *model.Transaction, logger *zap.Logger) {
	chainConfig, err := config.CtcGet(ctx).Chains.Get(transaction.ChainIDFrom)
	if err != nil {
		logger.Error("undefined chain config")
		return
	}

	var minedAt time.Time

	switch chainConfig.Kind {
	case string(model.ChainTypeEvm):
		poolByChain := modules.Pool().Clients(chainConfig)

		if tryErr := poolByChain.Try(ctx, func(ctx context.Context, client *ethclient.Client) error {
			txR, err := client.TransactionReceipt(ctx, common.BytesToHash(transaction.HashSource))
			if err != nil {
				return err
			}

			header, err := client.HeaderByHash(ctx, txR.BlockHash)
			if err != nil {
				return err
			}

			minedAt = time.Unix(int64(header.Time), 0)
			return nil
		}); tryErr != nil {
			logger.Error("can't calculate transaction duration", zap.Error(err), zap.String("tx_hash_source", common.BytesToHash(transaction.HashSource).String()))
			return
		}
	case string(model.ChainTypeCosmos):
		grpcConn := modules.TerraConn(chainConfig)
		txClient := tx.NewServiceClient(grpcConn)

		request := &tx.GetTxRequest{Hash: hex.EncodeToString(transaction.HashSource)}
		txR, err := txClient.GetTx(ctx, request)
		if err != nil {
			logger.Error("err get tx from chain", zap.Error(err), zap.String("tx_hash_source", common.BytesToHash(transaction.HashSource).String()))
			return
		}

		req := tmservice.GetBlockByHeightRequest{
			Height: txR.TxResponse.Height,
		}
		svcBlock := tmservice.NewServiceClient(grpcConn)
		resp, err := svcBlock.GetBlockByHeight(ctx, &req)
		if err != nil {
			logger.Error("can't calculate transaction duration", zap.Error(err), zap.String("tx_hash_source", common.BytesToHash(transaction.HashSource).String()))
			return
		}
		minedAt = resp.Block.Header.Time
	default:
		logger.Error(fmt.Sprintf("wrong chain type, received: %s", chainConfig.Kind))
		return
	}
	transaction.SourceTransactionMinedAt = &minedAt
	since := transaction.CreatedAt.Sub(minedAt)

	cLogger := logger.With(
		zap.Float64("processing_delay", math.Round(since.Seconds())),
		zap.Time("source_tx_mined_at", minedAt),
	)
	if since.Seconds() < maxProcessingDelay {
		cLogger.Debug("normal source tx duration")
	} else {
		cLogger.Error("long source tx duration")
	}

	if err := repository.Update(ctx, transaction); err != nil {
		logger.Error("err update source_tx_mined_at", zap.Error(err))
	}
}
