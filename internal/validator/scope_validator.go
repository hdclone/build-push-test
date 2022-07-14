package validator

import (
	"context"
)

type ValidatorsScope []*Validator

type Results map[string][]string

func (r Results) Valid() bool {
	return len(r) == 0
}

func Scope(validators ...*Validator) *ValidatorsScope {
	scope := ValidatorsScope(validators)
	return &scope
}

func (s *ValidatorsScope) Validate(ctx context.Context) (*Results, error) {
	results := &Results{}
	for _, validator := range *s {
		result, err := validator.Validate(ctx)
		if err != nil {
			return results, err
		} else if result != nil {
			if _, ok := (*results)[result.field]; ok {
				(*results)[result.field] = append((*results)[result.field], result.messages...)
			} else {
				(*results)[result.field] = result.messages
			}
		}
	}
	return results, nil
}
