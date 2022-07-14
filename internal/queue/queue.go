package queue

import (
	"broadcaster/internal/config"
	"broadcaster/internal/logging"
	"broadcaster/internal/model"
	"broadcaster/internal/repository/transactions"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ErrorWithStackTrace interface {
	StackTrace() errors.StackTrace
}

type Sender func(context.Context, *model.Transaction) error

type Queue struct {
	chain          *config.ChainConfig
	transactionIDs chan model.ID
	logger         *zap.Logger
	size           uint
	shutdown       context.CancelFunc
}

func New(logger *zap.Logger, size uint) *Queue {
	return &Queue{
		size:   size,
		logger: logger,
	}
}

func (queue *Queue) Start(ctx context.Context, chain *config.ChainConfig, sender Sender) {
	ctx, queue.shutdown = context.WithCancel(ctx)
	queue.logger.Info("start worker")
	queue.chain = chain
	queue.transactionIDs = make(chan model.ID, queue.size)
	go queue.worker(ctx, sender)
}

func (queue *Queue) Shutdown() error {
	if queue.shutdown != nil {
		queue.logger.Info("shutdown worker")
		queue.shutdown()
	}
	return nil
}

func (queue *Queue) Enqueue(transaction *model.Transaction) error {
	queue.logger.With(zap.String("transaction", string(transaction.ID)), zap.String("tx_hash_source", common.BytesToHash(transaction.HashSource).String())).Info("enqueue transaction")
	queue.transactionIDs <- transaction.ID
	return nil
}

func (queue *Queue) worker(ctx context.Context, sender func(context.Context, *model.Transaction) error) {
	defer close(queue.transactionIDs)
	for {
		select {
		case <-ctx.Done():
			return
		case transactionId, ok := <-queue.transactionIDs:
			if ok {
				var (
					transaction    model.Transaction
					logger         = queue.logger.With(zap.String("transaction", string(transactionId)))
					transactionCtx = logging.CtxSet(ctx, logger)
				)
				logger.Info("pickup transaction")
				if _, err := transactions.Query(transactionCtx).Apply(transactions.ID(transactionId)).Exec(transactionCtx, &transaction); err != nil {
					logger.Error(err.Error())
				} else {
					if err = sender(logging.CtxSet(ctx, logger.With(zap.String("tx_hash_source", common.BytesToHash(transaction.HashSource).String()))), &transaction); err != nil {
						logger.Error(err.Error())
					}
				}
			} else {
				return
			}
		}
	}
}
