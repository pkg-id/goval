package goval

import (
	"context"
	"time"
)

// TimeValidator is a FunctionValidator that validates time.Time.
type TimeValidator FunctionValidator[time.Time]

// Time returns a TimeValidator with no rules.
func Time() TimeValidator {
	return NopFunctionValidator[time.Time]
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

// When adds validation logic to the chain based on a condition.
//
// If the specified predicate returns true for an input value of type time.Time, the
// result of the chainer function is added to the chain, and the input value is
// validated using the new chain. If the predicate returns false, the original
// chain is returned without modification, and the input value is not validated.
func (f TimeValidator) When(predicate func(value time.Time) bool, chainer func(chain TimeValidator) TimeValidator) TimeValidator {
	return func(ctx context.Context, val time.Time) error {
		if predicate(val) {
			return chainer(f).Validate(ctx, val)
		}
		return f.Validate(ctx, val)
	}
}
