package rules

import (
	"broadcaster/internal/validator"
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun"
)

type ExistsScopeRule struct {
	*bun.SelectQuery
}

func (u ExistsScopeRule) Validate(ctx context.Context, _ interface{}) error {
	if result, err := u.Exec(ctx); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	} else if count, _ := result.RowsAffected(); count == 0 {
		return validator.Message("not found")
	}
	return nil
}
