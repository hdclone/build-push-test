package database

import (
	"context"

	"github.com/uptrace/bun"
)

type dbContextKey struct {
}

var dbKey = dbContextKey{}

func CtxGet(ctx context.Context) *bun.DB {
	if database, ok := ctx.Value(dbKey).(*bun.DB); ok {
		return database
	}
	return nil
}

func CtxSet(ctx context.Context, database *bun.DB) context.Context {
	return context.WithValue(ctx, dbKey, database)
}
