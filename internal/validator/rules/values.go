package rules

import (
	"broadcaster/internal/validator"
	"context"
)

type ValuesRule struct {
	values []interface{}
}

func (e *ValuesRule) Validate(_ context.Context, value interface{}) error {
	for _, checkValue := range e.values {
		if checkValue == value {
			return nil
		}
	}
	return validator.Message("not allowed value")
}

//func InValues(values []interface{}) *ValuesRule {
//	return &ValuesRule{values: values}
//}
