package goval

import (
	"context"
	"github.com/pkg-id/goval/funcs"
	"strings"
)

// StringValidator is a FunctionValidator that validates string.
type StringValidator FunctionValidator[string]

// String returns a StringValidator with no rules.
func String() StringValidator {
	return NopFunctionValidator[string]()
}

// Build builds the validator chain and attaches the value to it.
func (sv StringValidator) Build(value string) Validator {
	return validatorOf(sv, value)
}

// With attaches the next rule to the chain.
func (sv StringValidator) With(next StringValidator) StringValidator {
	return Chain(sv, next)
}

// Required ensures the string is not empty.
func (sv StringValidator) Required() StringValidator {
	return sv.With(func(ctx context.Context, value string) error {
		if value == "" {
			return NewRuleError(StringRequired, value)
		}
		return nil
	})
}

// Min ensures the length of the string is not less than the given length.
func (sv StringValidator) Min(length int) StringValidator {
	return sv.With(func(ctx context.Context, value string) error {
		if len(value) < length {
			return NewRuleError(StringMin, value, length)
		}
		return nil
	})
}

// Max ensures the length of the string is not greater than the given length.
func (sv StringValidator) Max(length int) StringValidator {
	return sv.With(func(ctx context.Context, value string) error {
		if len(value) > length {
			return NewRuleError(StringMax, value, length)
		}
		return nil
	})
}

// Match ensures the string matches the given pattern.
func (sv StringValidator) Match(pattern Pattern) StringValidator {
	return sv.With(func(ctx context.Context, value string) error {
		exp := pattern.RegExp()
		if !exp.MatchString(value) {
			return NewRuleError(StringMatch, value, exp.String())
		}
		return nil
	})
}

// In ensures that the provided string is one of the specified options.
// This validation is case-sensitive, use InFold to perform a case-insensitive In validation.
func (sv StringValidator) In(options ...string) StringValidator {
	return sv.With(func(ctx context.Context, value string) error {
		ok := funcs.Contains(options, func(opt string) bool { return value == opt })
		if !ok {
			return NewRuleError(StringIn, value, options)
		}
		return nil
	})
}

// InFold ensures that the provided string is one of the specified options with case-insensitivity.
func (sv StringValidator) InFold(options ...string) StringValidator {
	return sv.With(func(ctx context.Context, value string) error {
		ok := funcs.Contains(options, func(opt string) bool { return strings.EqualFold(value, opt) })
		if !ok {
			return NewRuleError(StringInFold, value, options)
		}
		return nil
	})
}
