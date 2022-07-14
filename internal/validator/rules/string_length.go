package rules

import (
	"broadcaster/internal/validator"
	"context"
)

type StringLengthRule struct {
	min int
	max int
	len int
}

func (p *StringLengthRule) Validate(_ context.Context, value interface{}) error {
	l := len(value.(string))
	if p.max != 0 && l > p.max {
		return validator.Message("is too long")
	}
	if p.min != 0 && l < p.min {
		return validator.Message("is too short")
	}
	if p.len != 0 && l != p.len {
		return validator.Message("is the wrong length")
	}
	return nil
}
