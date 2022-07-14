package rules

import (
	"broadcaster/internal/validator"
	"context"
	"regexp"
)

type RegExpRule struct {
	re *regexp.Regexp
}

func (e RegExpRule) Validate(_ context.Context, value interface{}) error {
	if !e.re.MatchString(value.(string)) {
		return validator.Message("invalid format")
	}
	return nil
}

//func IsRegExp(expression string) *RegExpRule {
//	return &RegExpRule{regexp.MustCompile(expression)}
//}
