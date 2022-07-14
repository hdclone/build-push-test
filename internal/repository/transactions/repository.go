package transactions

import (
	"broadcaster/internal/database"
	"broadcaster/internal/model"
	"context"
	"github.com/uptrace/bun"
)

func Query(ctx context.Context) *bun.SelectQuery {
	return database.CtxGet(ctx).NewSelect().Model((*model.Transaction)(nil))
}
