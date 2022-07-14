package middlewares

import (
	"broadcaster/internal/logging"
	"broadcaster/internal/response"
	"errors"
	"fmt"
	"net/http"
)

func Catch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			recovered := recover()
			if recovered != nil {
				switch err := recovered.(type) {
				case string:
					response.InternalServerError().WithError(errors.New(err)).Write(w, r)
					logging.CtxGet(r.Context()).Error(err)
				case response.ErrorResponse:
					err.Write(w, r)
					if err.Code() == http.StatusInternalServerError {
						logging.CtxGet(r.Context()).Error(err.Error().Error())
					}
				case *response.ErrorResponse:
					err.Write(w, r)
					if err.Code() == http.StatusInternalServerError {
						logging.CtxGet(r.Context()).Error(err.Error().Error())
					}
				default:
					response.InternalServerError().Write(w, r)
					logging.CtxGet(r.Context()).Error(fmt.Sprint(recovered))
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
