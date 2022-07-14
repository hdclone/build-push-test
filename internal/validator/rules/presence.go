package rules

import (
	"broadcaster/internal/validator"
	"context"
	"reflect"
)

type PresenceRule struct {
}

func (p *PresenceRule) Validate(_ context.Context, value interface{}) error {
	v := reflect.ValueOf(value)
	invalid := false

	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		if v.Len() == 0 {
			invalid = true
		}
	case reflect.Bool:
		if !v.Bool() {
			invalid = true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			invalid = true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() == 0 {
			invalid = true
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0 {
			invalid = true
		}
	case reflect.Invalid:
		invalid = true
	case reflect.Interface, reflect.Ptr:
		invalid = v.IsNil()
	}

	if invalid {
		return validator.Message("can't be blank")
	}
	return nil
}

func IsPresence() *PresenceRule {
	return &PresenceRule{}
}
