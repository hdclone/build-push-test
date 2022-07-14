package modules

import (
	"broadcaster/internal/queue"
	"fmt"
	"go.uber.org/zap"
)

func Queue(chainId int64) *queue.Queue {
	return Register(fmt.Sprintf("queue-%d", chainId), func(s string) (Module, error) {
		return queue.New(Logger().With(zap.String("module", "queue"), zap.Int64("chain", chainId)), Config().Queue.Size), nil
	}).(*queue.Queue)
}
