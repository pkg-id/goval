package goval

import (
	"context"
	"regexp"
)

// Validator is an interface for all validators. It provides a contract for grouping different kinds of validators.
// Each validator, such as StringValidator, delays the execution of its rules until it is invoked via the Validate method.
type Validator interface {
	// Validate starts the execution of the Validator.
	// When this method is called, the Validator will execute all the rules in the chain,
	// and return the first error encountered.
	//
	// By default, if the error is a RuleError, it will be translated by DefaultErrorTranslator.
	// To customize the error translation, use the SetErrorTranslator method.
	Validate(ctx context.Context) error
}

// ValidatorFunc is an adapter for creating an implementation of Validator by using an ordinary function.
type ValidatorFunc func(ctx context.Context) error

// Validate implements the Validator interface by invoking itself.
func (f ValidatorFunc) Validate(ctx context.Context) error { return f(ctx) }

// validatorOf is a helper function that creates a Validator from a FunctionValidator and a value.
func validatorOf[T any](fn func(ctx context.Context, value T) error, value T) Validator {
	return ValidatorFunc(func(ctx context.Context) error {
		err := fn(ctx, value)
		return translateValidatorError(ctx, err)
	})
}

// translateValidatorError translates the error if it is a RuleError.
// Otherwise, it returns the error as is.
func translateValidatorError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	switch et := err.(type) {
	default:
		return err
	case *RuleError:
		return globalErrorTranslator.Translate(ctx, et)
	}
}

// NopFunctionValidator does nothing and always returns nil.
// It is useful as the first rule in the validator chain.
// Important to note that the Chain function requires at least two functions, by using NopFunctionValidator
// as the first rule, we ensure that the Chain function will always have at least two functions to chain.
func NopFunctionValidator[T any]() func(context.Context, T) error {
	return func(ctx context.Context, t T) error { return nil }
}

// Builder is an interface that defines the Validator from the ValueValidator.
// The Builder interface is used to construct the validator that can be used to start the validation process.
// This Builder acts like a factory for Validator and also as the input supplier for the ValueValidator.
type Builder[T any] interface {
	// Build returns a validator that can be used to start the validation process.
	Build(value T) Validator
}

// BuilderFunc is an adapter for creating an implementation of Builder by using an ordinary function.
type BuilderFunc[T any] func(value T) Validator

// Build implements the Builder interface by invoking itself.
func (f BuilderFunc[T]) Build(value T) Validator { return f(value) }

// FunctionValidator is a signature for a function that validates any value.
// The context is used to pass additional information to the validator, such as the locale.
type FunctionValidator[T any] func(ctx context.Context, value T) error

// FunctionValidatorConstraint is a set of functions signatures that treats the FunctionValidator.
type FunctionValidatorConstraint[T any] interface {
	~func(ctx context.Context, value T) error
}

// Chain creates a new function that chains the execution of two given functions into a single function.
// Here's an example: suppose we have two functions:
//
//	var f func(ctx context.Context, value T) error
//	var g func(ctx context.Context, value T) error
//
// We want to combine `f` and `g` into a single function, but without executing them immediately. In other words, we
// want to delay the execution of `f` and `g` until the new function is executed.
// Let's call the new function `h`. When `h` is executed, `f` will be executed first. If `f` executes without error,
// then `g` will be executed next. If `h` returns any error it will be an error that returned either `f` or `g`.
func Chain[T any, Func FunctionValidatorConstraint[T]](f, g Func) Func {
	return func(ctx context.Context, value T) error {
		return execChain(ctx, value, f, g)
	}
}

// execChain executes the given functions in the order they are given.
// If any of the functions returns an error, the execution will be stopped and the error will be returned.
func execChain[T any, Func FunctionValidatorConstraint[T]](ctx context.Context, value T, functions ...Func) error {
	for _, fn := range functions {
		if err := fn(ctx, value); err != nil {
			return err
		}
	}
	return nil
}

// Named creates a new validator that returns KeyErrors if actual validator returns an error.
func Named[T any, B Builder[T]](name string, value T, validator B) Validator {
	return ValidatorFunc(func(ctx context.Context) error {
		if err := validator.Build(value).Validate(ctx); err != nil {
			return NewKeyError(name, err)
		}
		return nil
	})
}

// Each creates a slice validator that validates each element in the slice.
func Each[T any, V []T](fn func(each T) Validator) Builder[V] {
	return BuilderFunc[V](func(values V) Validator {
		return ValidatorFunc(func(ctx context.Context) error {
			return each(ctx, values, fn)
		})
	})
}

// each executes the given function for each element in the slice.
func each[T any, V []T](ctx context.Context, values V, fn func(value T) Validator) error {
	errs := make(Errors, 0)
	for _, value := range values {
		err := fn(value).Validate(ctx)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

// Execute executes the given validators and collects the errors into a single error
func Execute(ctx context.Context, validators ...Validator) error {
	errs := make(Errors, 0)
	for _, validator := range validators {
		if validator != nil {
			err := validator.Validate(ctx)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

// Pattern is an interface that defines the regular expression pattern.
type Pattern interface {
	// RegExp returns the compiled regular expression.
	RegExp() *regexp.Regexp
}
