package v1

import (
	"broadcaster/internal/response"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	response.Ok().Write(w, r)
}
