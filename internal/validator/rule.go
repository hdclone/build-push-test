package validator

import "context"

type Rule interface {
	Validate(ctx context.Context, value interface{}) error
}

type Message string

func (m Message) Error() string {
	return string(m)
}
