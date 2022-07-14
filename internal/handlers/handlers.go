package handlers

import (
	"broadcaster/internal/response"
	"github.com/gorilla/mux"
	"net/http"
)

func Default(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.NotFoundError().Write(w, r)
	})
	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.MethodNotAllowed().Write(w, r)
	})
}
