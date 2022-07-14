package middlewares

import (
	"broadcaster/internal/logging"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Response struct {
	http.ResponseWriter
	Status int
}

func (r *Response) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		logger := logging.CtxGet(r.Context())
		response := &Response{w, http.StatusOK}
		defer func(t time.Time) {
			logger.With(
				zap.Int("code", response.Status),
				zap.Duration("elapsed", time.Since(t)),
			).Debug("Complete request")
		}(startTime)
		logger.Debug("Started request")
		next.ServeHTTP(response, r)
	})
}
