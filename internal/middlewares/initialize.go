package middlewares

import (
	"broadcaster/internal/config"
	"broadcaster/internal/database"
	"broadcaster/internal/logging"
	"broadcaster/internal/modules"
	"broadcaster/internal/pool"
	"broadcaster/internal/server"
	"broadcaster/internal/variables"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"net"
	"net/http"
)

func Initialize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var (
			requestId uuid.UUID
			err       error
		)

		if value := r.Header.Get("X-Request-ID"); len(value) > 0 {
			if requestId, err = uuid.FromString(value); err != nil {
				requestId, _ = uuid.DefaultGenerator.NewV4()
			}
		} else {
			requestId, _ = uuid.DefaultGenerator.NewV4()
		}

		w.Header().Set("X-Service", variables.Service("Service"))
		w.Header().Set("X-Service-Version", variables.Version)
		w.Header().Set("X-Request-ID", requestId.String())

		var realIP string
		if realIP = r.Header.Get("CF-Connecting-IP"); len(realIP) == 0 {
			if realIP = r.Header.Get("X-Real-IP"); len(realIP) == 0 {
				realIP, _, _ = net.SplitHostPort(r.RemoteAddr)
			}
		}

		ctx = server.CtxSetRequestID(ctx, &requestId)
		ctx = config.CtxSet(ctx, modules.Config())
		ctx = logging.CtxSet(ctx, modules.Logger().With(
			zap.String("remote-ip", realIP),
			zap.String("request-id", requestId.String()),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path)))

		ctx = database.CtxSet(ctx, modules.Database())
		ctx = pool.CtxSet(ctx, modules.Pool())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
