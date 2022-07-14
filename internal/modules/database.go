package modules

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Database() *bun.DB {
	return Register("database", func(s string) (Module, error) {
		sqlDb, err := sql.Open("pgx", Config().Store.DSN)
		if err != nil {
			return nil, err
		}
		db := bun.NewDB(sqlDb, pgdialect.New())
		//db.AddQueryHook(&database.LoggerHook{
		//	SlowThreshold: 200 * time.Millisecond,
		//	QueryLevel:    zap.DebugLevel,
		//	SlowLevel:     zap.WarnLevel,
		//	ErrorLevel:    zap.ErrorLevel,
		//})
		return db, nil
	}).(*bun.DB)
}
