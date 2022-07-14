package logging

import (
	"context"

	"go.uber.org/zap"
)

type loggerContextKey struct {
}

var logKey = loggerContextKey{}

func CtxGet(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(logKey).(*zap.Logger); ok {
		return logger
	}
	return nil
}

func CtxSet(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, logKey, logger)
}
