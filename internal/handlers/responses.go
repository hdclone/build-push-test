package handlers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type Details = map[string]interface{}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Write(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	_ = json.NewEncoder(w).Encode(e)
}

func (e Error) WithDetails(details interface{}) Error {
	e.Details = details
	return e
}

func Custom(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NotFoundError() *Error {
	return Custom(http.StatusNotFound, "not found")
}
