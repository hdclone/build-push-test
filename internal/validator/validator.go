package validator

import (
	"context"
)

type Validator struct {
	field string
	value interface{}
	rules []Rule
}

type Result struct {
	field    string
	messages []string
}

func (v *Validator) Validate(ctx context.Context) (*Result, error) {
	result := &Result{
		field:    v.field,
		messages: []string{},
	}
	for _, rule := range v.rules {
		validateResult := rule.Validate(ctx, v.value)
		if message, ok := validateResult.(Message); ok {
			result.messages = append(result.messages, string(message))
		} else if validateResult != nil {
			return nil, validateResult
		}
	}
	if len(result.messages) > 0 {
		return result, nil
	}
	return nil, nil
}

func (v *Validator) Rules(rules ...Rule) *Validator {
	nv := *v
	nv.rules = append(nv.rules, rules...)
	return &nv
}

func Validate(field string, value interface{}) *Validator {
	return &Validator{field: field, value: value}
}
