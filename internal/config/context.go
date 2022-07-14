package config

import (
	"context"
)

type configContextKey struct {
}

var cfk = configContextKey{}

func CtcGet(ctx context.Context) *Config {
	return ctx.Value(cfk).(*Config)
}

func CtxSet(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, cfk, cfg)
}
