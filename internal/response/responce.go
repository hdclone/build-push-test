package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	code    int
	message string
	error   error
	data    interface{}
}

type Details = map[string]interface{}

func (e ErrorResponse) Message() string {
	return e.message
}

func (e ErrorResponse) Error() error {
	return e.error
}

func (e ErrorResponse) Code() int {
	return e.code
}

func (e ErrorResponse) Data() interface{} {
	return e.data
}

func (e *ErrorResponse) WithCode(code int) *ErrorResponse {
	e.code = code
	return e
}

func (e *ErrorResponse) WithError(err error) *ErrorResponse {
	e.error = err
	return e
}

func (e *ErrorResponse) WithMessage(message string, args ...interface{}) *ErrorResponse {
	e.message = fmt.Sprintf(message, args...)
	return e
}

func (e *ErrorResponse) WithData(data interface{}) *ErrorResponse {
	e.data = data
	return e
}

func (e ErrorResponse) Write(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.code)
	_ = json.NewEncoder(w).Encode(struct {
		Code    int         `json:"code"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"details,omitempty"`
	}{
		Code:    e.code,
		Message: e.message,
		Data:    e.data,
	})
}

func New() *ErrorResponse {
	return &ErrorResponse{}
}

func InternalServerError() *ErrorResponse {
	return New().WithCode(http.StatusInternalServerError).WithMessage("internal server error")
}

func NotFoundError() *ErrorResponse {
	return New().WithCode(http.StatusNotFound).WithMessage("not found")
}

func MethodNotAllowed() *ErrorResponse {
	return New().WithCode(http.StatusMethodNotAllowed).WithMessage("not allowed")
}

func Ok() *ErrorResponse {
	return New().WithCode(http.StatusOK)
}

func BadRequest() *ErrorResponse {
	return New().WithCode(http.StatusBadRequest)
}

func UnprocessableEntity() *ErrorResponse {
	return New().WithCode(http.StatusUnprocessableEntity)
}
