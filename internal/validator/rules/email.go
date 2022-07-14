package rules

import (
	"broadcaster/internal/validator"
	"context"
	"net/mail"
)

type EmailRule struct {
}

func (e EmailRule) Validate(_ context.Context, value interface{}) error {
	if _, err := mail.ParseAddress(value.(string)); err != nil {
		return validator.Message("invalid email")
	}
	return nil
}
