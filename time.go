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

// Validate executes the validation rules immediately.
func (f TimeValidator) Validate(ctx context.Context, value time.Time) error {
	return validatorOf(f, value).Validate(ctx)
}

// With attaches the next rule to the chain.
func (f TimeValidator) With(next TimeValidator) TimeValidator {
	return Chain(f, next)
}

// Required ensures the time is not zero.
func (f TimeValidator) Required() TimeValidator {
	return f.With(func(ctx context.Context, value time.Time) error {
		if value.IsZero() {
			return NewRuleError(TimeRequired)
		}
		return nil
	})
}

// Min ensures the time is after min.
func (f TimeValidator) Min(min time.Time) TimeValidator {
	return f.With(func(ctx context.Context, value time.Time) error {
		if value.Before(min) {
			return NewRuleError(TimeMin, min)
		}
		return nil
	})
}

// Max ensures the time is before max.
func (f TimeValidator) Max(max time.Time) TimeValidator {
	return f.With(func(ctx context.Context, value time.Time) error {
		if value.After(max) {
			return NewRuleError(TimeMax, max)
		}
		return nil
	})
}
