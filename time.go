package goval

import (
	"context"
	"time"
)

// TimeValidator is a validator for time.Time type.
type TimeValidator FunctionValidator[time.Time]

// Time is TimeValidator constructor. This function is used to initialize
// the rules chain. Since, it will be a first rule in the chain, it not validates anything.
func Time() TimeValidator {
	return NopFunctionValidator[time.Time]()
}

// Build attaches the value so the rules chain can consume it as an input that need to be validated.
func (tv TimeValidator) Build(value time.Time) Validator {
	return validatorOf(tv, value)
}

// With added TimeValidator to the rules chain.
func (tv TimeValidator) With(next TimeValidator) TimeValidator {
	return Chain(tv, next)
}

// Required ensures the time is not zero time instant, January 1, year 1, 00:00:00 UTC.
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
