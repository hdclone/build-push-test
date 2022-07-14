package pool

import "context"

type poolContextKey struct{}

var poolKey = poolContextKey{}

func CtcGet(ctx context.Context) *Pool {
	return ctx.Value(poolKey).(*Pool)
}

func CtxSet(ctx context.Context, pool *Pool) context.Context {
	return context.WithValue(ctx, poolKey, pool)
}
