package goval

import (
	"context"
	"github.com/pkg-id/goval/funcs"
)

// SliceValidator is a FunctionValidator that validates slices.
type SliceValidator[T any, V []T] FunctionValidator[V]

// Slice returns a SliceValidator with no rules.
// T is the type of the slice elements, V is the type of the slice.
func Slice[T any, V []T]() SliceValidator[T, V] {
	return NopFunctionValidator[V]()
}

// Validate executes the validation rules immediately.
func (f SliceValidator[T, V]) Validate(ctx context.Context, values V) error {
	return validatorOf(f, values).Validate(ctx)
}

// With attaches the next rule to the chain.
func (f SliceValidator[T, V]) With(next SliceValidator[T, V]) SliceValidator[T, V] {
	return Chain(f, next)
}

// Required ensures the slice is not empty.
func (f SliceValidator[T, V]) Required() SliceValidator[T, V] {
	return f.With(func(ctx context.Context, values V) error {
		if len(values) == 0 {
			return NewRuleError(SliceRequired)
		}
		return nil
	})
}

// Min ensures the length of the slice is not less than the given min.
func (f SliceValidator[T, V]) Min(min int) SliceValidator[T, V] {
	return f.With(func(ctx context.Context, values V) error {
		if len(values) < min {
			return NewRuleError(SliceMin, min)
		}
		return nil
	})
}

// Max ensures the length of the slice is not greater than the given max.
func (f SliceValidator[T, V]) Max(max int) SliceValidator[T, V] {
	return f.With(func(ctx context.Context, values V) error {
		if len(values) > max {
			return NewRuleError(SliceMax, max)
		}
		return nil
	})
}

// Each ensures each element of the slice is satisfied by the given validator.
func (f SliceValidator[T, V]) Each(validator RuleValidator[T]) SliceValidator[T, V] {
	return func(ctx context.Context, values V) error {
		validators := funcs.Map(values, RuleValidatorToValidatorFactory(validator))
		return execute(ctx, validators)
	}
}
