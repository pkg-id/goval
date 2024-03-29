package goval

import (
	"context"
	"github.com/pkg-id/goval/funcs"
)

// MapValidator is a FunctionValidator that validates map[K]V.
type MapValidator[K comparable, V any] FunctionValidator[map[K]V]

// Map returns a MapValidator with no rules.
func Map[K comparable, V any]() MapValidator[K, V] {
	return NopFunctionValidator[map[K]V]
}

// Validate executes the validation rules immediately.
func (f MapValidator[K, V]) Validate(ctx context.Context, values map[K]V) error {
	return validatorOf(f, values).Validate(ctx)
}

// With attaches the next rule to the chain.
func (f MapValidator[K, V]) With(next MapValidator[K, V]) MapValidator[K, V] {
	return Chain(f, next)
}

// Required ensures the length is not zero.
func (f MapValidator[K, V]) Required() MapValidator[K, V] {
	return f.With(func(ctx context.Context, values map[K]V) error {
		if len(values) == 0 {
			return NewRuleError(MapRequired)
		}
		return nil
	})
}

// Min ensures the length is not less than the given min.
func (f MapValidator[K, V]) Min(min int) MapValidator[K, V] {
	return f.With(func(ctx context.Context, values map[K]V) error {
		if len(values) < min {
			return NewRuleError(MapMin, min)
		}
		return nil
	})
}

// Max ensures the length is not greater than the given max.
func (f MapValidator[K, V]) Max(max int) MapValidator[K, V] {
	return f.With(func(ctx context.Context, values map[K]V) error {
		if len(values) > max {
			return NewRuleError(MapMax, max)
		}
		return nil
	})
}

// Each ensures each element of the map is satisfied by the given validator.
func (f MapValidator[K, V]) Each(validator RuleValidator[V]) MapValidator[K, V] {
	return f.With(func(ctx context.Context, values map[K]V) error {
		validators := funcs.Map(funcs.Values(values), RuleValidatorToValidatorFactory(validator))
		return execute(ctx, validators)
	})
}
