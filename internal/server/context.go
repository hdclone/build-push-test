package server

import (
	"context"

	"github.com/gofrs/uuid"
)

type requestContextKey struct {
}

var requestKey = requestContextKey{}

func CtxGetRequestID(ctx context.Context) *uuid.UUID {
	if requestId, ok := ctx.Value(requestKey).(*uuid.UUID); ok {
		return requestId
	}
	return nil
}

func CtxSetRequestID(ctx context.Context, requestId *uuid.UUID) context.Context {
	return context.WithValue(ctx, requestKey, requestId)
}
