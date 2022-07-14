package middlewares

import (
	"github.com/rs/cors"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	c := cors.AllowAll()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.ServeHTTP(w, r, func(writer http.ResponseWriter, request *http.Request) {
			next.ServeHTTP(writer, request)
		})
	})
}
