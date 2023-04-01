package goval

import (
	"context"
	"github.com/pkg-id/goval/constraints"
	"github.com/pkg-id/goval/funcs"
)

// NumberConstraint is set types that treats as numbers.
type NumberConstraint interface {
	constraints.Integer | constraints.Float
}

// NumberValidator is a validator that validates numbers.
type NumberValidator[T NumberConstraint] FunctionValidator[T]

// Number returns a NumberValidator with no rules.
func Number[T NumberConstraint]() NumberValidator[T] {
	return NopFunctionValidator[T]
}

// Validate executes the validation rules immediately.
func (f NumberValidator[T]) Validate(ctx context.Context, value T) error {
	return validatorOf(f, value).Validate(ctx)
}

// With attaches the next rule to the chain.
func (f NumberValidator[T]) With(next NumberValidator[T]) NumberValidator[T] {
	return Chain(f, next)
}

// Required ensures the number is not zero.
func (f NumberValidator[T]) Required() NumberValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		var zero T
		if value == zero {
			return NewRuleError(NumberRequired)
		}
		return nil
	})
}

// Min ensures the number is not less than the given min.
func (f NumberValidator[T]) Min(min T) NumberValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		if value < min {
			return NewRuleError(NumberMin, min)
		}
		return nil
	})
}

// Max ensures the number is not greater than the given max.
func (f NumberValidator[T]) Max(max T) NumberValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		if value > max {
			return NewRuleError(NumberMax, max)
		}
		return nil
	})
}

// In ensures that the provided number is one of the specified options.
func (f NumberValidator[T]) In(options ...T) NumberValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		ok := funcs.Contains(options, func(opt T) bool { return opt == value })
		if !ok {
			return NewRuleError(NumberIn, options)
		}
		return nil
	})
}

// When adds validation logic to the chain based on a condition.
//
// If the specified predicate returns true for an input value of type T, the
// result of the chainer function is added to the chain, and the input value is
// validated using the new chain. If the predicate returns false, the original
// chain is returned without modification, and the input value is not validated.
//
// The chainer function takes a NumberValidator[T] instance as input and returns
// a new NumberValidator[T] instance that includes additional validation logic.
//
// When returns a new NumberValidator[T] instance that can be used to validate
// values of type T, with the added validation logic from the chainer function.
func (f NumberValidator[T]) When(p Predicate[T], l Linker[T, NumberValidator[T]]) NumberValidator[T] {
	return whenLinker(f, p, l)
}
