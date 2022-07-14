package rules

import (
	"broadcaster/internal/validator"
	"context"
)

type CustomCallback func() error

type CustomRule struct {
	callback CustomCallback
}

func IsCustom(callback CustomCallback) validator.Rule {
	return &CustomRule{callback: callback}
}

func (c *CustomRule) Validate(_ context.Context, _ interface{}) error {
	return c.callback()
}
