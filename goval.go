package goval

import (
	"context"
	"encoding/json"
	"fmt"
)

// Error is a type that represents a simple string message as an error.
// Unlike the error returned by errors.New, this error does not need to implement json.Marshaler.
// This is because the string literal itself is already a valid JSON format. This error is created from a string type,
// which means that it can be initialized as a const and when this error is compared, it will be compared by the
// string value within it. For example:
//
//	const err1 = Error("abc")
//	const err2 = Error("abc")
//	const err3 = Error("xyz")
//	fmt.Println(err1 == err2, err1 == err3) // true false
type Error string

// Error implements the built-in error interface.
func (e Error) Error() string { return string(e) }

// Errorf is a helper function for creating an Error with a format.
func Errorf(format string, args ...any) error {
	return Error(fmt.Sprintf(format, args...))
}

type KeyError struct {
	Key string `json:"key"`
	Err error  `json:"err"`
}

type auxKeyError KeyError

func NewKeyError(key string, err error) *KeyError {
	return &KeyError{
		Key: key,
		Err: err,
	}
}

func (k *KeyError) Error() string {
	return k.String()
}

func (k *KeyError) String() string {
	b, err := k.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (k *KeyError) MarshalJSON() ([]byte, error) {
	aux := auxKeyError(*k)
	if _, ok := k.Err.(json.Marshaler); !ok {
		aux.Err = Error(k.Err.Error())
	}

	b, err := json.Marshal(aux)
	if err != nil {
		return nil, fmt.Errorf("goval: KeyError.MarshalJSON: %w", err)
	}

	return b, nil
}

// Errors is a type for collecting multiple errors and bundling them into a single error.
type Errors struct {
	errs []error
}

// NewErrors creates a new error collector.
func NewErrors() *Errors {
	return &Errors{errs: make([]error, 0)}
}

// Append appends the given error to the internal error collector if it is not nil.
func (e *Errors) Append(err error) {
	if err != nil {
		e.errs = append(e.errs, err)
	}
}

// Err returns a nil error if no errors are found in the internal error collector.
// Otherwise, it returns the bundled errors.
func (e *Errors) Err() error {
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

func (e *Errors) Errs() []error {
	return e.errs
}

// Error implements the built-in error interface.
// This method returns the same value as the String method.
func (e *Errors) Error() string { return e.String() }

// String implements the built-in fmt.Stringer interface.
// This method returns the errors in JSON format.
func (e *Errors) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// MarshalJSON implements the built-in json.Marshaler interface for errors.
func (e *Errors) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(e.errs)
	if err != nil {
		return nil, fmt.Errorf("goval: Errors.MarshalJSON: %w", err)
	}
	return b, nil
}

// Validator is an interface for all validators. It provides a contract for grouping different kinds of validators.
// Each validator, such as StringValidator, delays the execution of its rules until it is
// invoked via the Validate method.
type Validator interface {
	// Validate starts the execution of the Validator.
	Validate(ctx context.Context) error
}

// ValidatorFunc is an adapter for creating an implementation of Validator by using an ordinary function.
type ValidatorFunc func(ctx context.Context) error

// Validate implements the Validator interface by invoking itself.
func (f ValidatorFunc) Validate(ctx context.Context) error { return f(ctx) }

// validatorOf is a helper function that creates a Validator from a value validator function.
func validatorOf[T any](fn func(ctx context.Context, value T) error, value T) Validator {
	return ValidatorFunc(func(ctx context.Context) error {
		return fn(ctx, value)
	})
}

// NopValueValidator does nothing and always returns nil. It's meant to be used as the first validator in a chain.
func NopValueValidator[T any](ctx context.Context, v T) error {
	return nil
}

// Builder is an interface that defines the Validator from the ValueValidator.
// The Builder interface is used to construct the validator that can be used to start the validation process.
// This Builder acts like a factory for Validator and also as the input supplier for the ValueValidator.
type Builder[T any] interface {
	// Build returns a validator that can be used to start the validation process.
	Build(value T) Validator
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
func Chain[T any](f, g func(ctx context.Context, value T) error) func(ctx context.Context, value T) error {
	return func(ctx context.Context, value T) error {
		if err := f(ctx, value); err != nil {
			return err
		}

		return g(ctx, value)
	}
}

func Named(name string, validator Validator) Validator {
	return ValidatorFunc(func(ctx context.Context) error {
		if err := validator.Validate(ctx); err != nil {
			return NewKeyError(name, err)
		}
		return nil
	})
}

// Execute executes the given validators and collects the errors into a single error
func Execute(ctx context.Context, validators ...Validator) error {
	errs := NewErrors()
	for _, validator := range validators {
		if validator != nil {
			errs.Append(validator.Validate(ctx))
		}
	}
	return errs.Err()
}
