package goval

import (
	"context"
	"time"
)

// TimeValidator is a FunctionValidator that validates time.Time.
type TimeValidator FunctionValidator[time.Time]

// Time returns a TimeValidator with no rules.
func Time() TimeValidator {
	return NopFunctionValidator[time.Time]()
}

// Build builds the validator chain and attaches the value to it.
func (tv TimeValidator) Build(value time.Time) Validator {
	return validatorOf(tv, value)
}

// With attaches the next rule to the chain.
func (tv TimeValidator) With(next TimeValidator) TimeValidator {
	return Chain(tv, next)
}

// Required ensures the time is not zero.
func (tv TimeValidator) Required() TimeValidator {
	return tv.With(func(ctx context.Context, value time.Time) error {
		if value.IsZero() {
			return NewRuleError(TimeRequired, value)
		}
		return nil
	})
}

// Min ensures the time is after min.
func (tv TimeValidator) Min(min time.Time) TimeValidator {
	return tv.With(func(ctx context.Context, value time.Time) error {
		if value.Before(min) {
			return NewRuleError(TimeMin, value, min)
		}
		return nil
	})
}

// Max ensures the time is before max.
func (tv TimeValidator) Max(max time.Time) TimeValidator {
	return tv.With(func(ctx context.Context, value time.Time) error {
		if value.After(max) {
			return NewRuleError(TimeMax, value, max)
		}
		return nil
	})
}
