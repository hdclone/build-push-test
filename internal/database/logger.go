package database

import (
	"broadcaster/internal/logging"
	"bytes"
	"context"
	"time"

	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerHook struct {
	SlowThreshold time.Duration
	QueryLevel    zapcore.Level
	SlowLevel     zapcore.Level
	ErrorLevel    zapcore.Level
}

func (v *LoggerHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {

	log := logging.CtxGet(ctx)
	//dur := now.Sub(event.StartTime)
	//
	//switch event.Err {
	//case nil, sql.ErrNoRows:
	//	if v.SlowThreshold > 0 && dur >= v.SlowThreshold {
	//		level = v.SlowLevel
	//	} else {
	//		level = v.QueryLevel
	//	}
	//default:
	//	level = v.ErrorLevel
	//}

	log = log.With(zap.String("elapsed", time.Since(event.StartTime).String()))

	log.Debug(event.Query)
}

func (v *LoggerHook) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context {
	return ctx
}

// taken from bun
func eventOperation(event *bun.QueryEvent) string {
	switch event.QueryAppender.(type) {
	case *bun.SelectQuery:
		return "SELECT"
	case *bun.InsertQuery:
		return "INSERT"
	case *bun.UpdateQuery:
		return "UPDATE"
	case *bun.DeleteQuery:
		return "DELETE"
	case *bun.CreateTableQuery:
		return "CREATE TABLE"
	case *bun.DropTableQuery:
		return "DROP TABLE"
	}
	return queryOperation([]byte(event.Query))
}

// taken from bun
func queryOperation(name []byte) string {
	if idx := bytes.IndexByte(name, ' '); idx > 0 {
		name = name[:idx]
	}
	if len(name) > 16 {
		name = name[:16]
	}
	return string(name)
}
