package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func Mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
