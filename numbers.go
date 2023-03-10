package goval

import (
	"context"
	"golang.org/x/exp/constraints"
)

// NumberConstraint is a generic constraint for type Integers and Floats.
type NumberConstraint interface {
	constraints.Integer | constraints.Float
}

// NumberValidator is a validator for NumberConstraint type.
type NumberValidator[T NumberConstraint] func(ctx context.Context, value T) error

// Number is NumberValidator constructor. This function is used to initialize
// the rules chain. Since, it will be a first rule in the chain, it not validates anything.
func Number[T NumberConstraint]() NumberValidator[T] {
	return NopValueValidator[T]
}

// WithValue attaches the value so the rules chain can consume it as an input that need to be validated.
func (nv NumberValidator[T]) WithValue(value T) Validator {
	return validatorOf(nv, value)
}

// Required ensures the number is not a zero value.
func (nv NumberValidator[T]) Required() NumberValidator[T] {
	return Chain[T](nv, func(ctx context.Context, value T) error {
		var zero T
		if value == zero {
			return Error("is required")
		}
		return nil
	})
}

// Min ensures the number is not less than the given min.
func (nv NumberValidator[T]) Min(min T) NumberValidator[T] {
	return Chain(nv, func(ctx context.Context, value T) error {
		if value < min {
			return Errorf("must be greater than %v", min)
		}
		return nil
	})
}

// Max ensures the number is not greater than the given max.
func (nv NumberValidator[T]) Max(max T) NumberValidator[T] {
	return Chain(nv, func(ctx context.Context, value T) error {
		if value > max {
			return Errorf("must be less than %v", max)
		}
		return nil
	})
}
