package response

import "net/http"

type Response struct {
	http.ResponseWriter
	Status int
}

func (r *Response) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
