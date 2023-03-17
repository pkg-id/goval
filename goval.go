package goval

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"sync"
)

// TextError is an error type for turning an ordinary string to an error.
type TextError string

// Error implements the built-in error.
func (t TextError) Error() string { return t.String() }

// String implements the fmt.Stringer.
func (t TextError) String() string { return string(t) }

// MarshalJSON implements the json.Marshaler.
func (t TextError) MarshalJSON() ([]byte, error) { return json.Marshal(t.String()) }

type ErrorTranslator interface {
	Translate(ctx context.Context, err *RuleError) error
}

type ErrorTranslatorFunc func(ctx context.Context, err *RuleError) error

func (f ErrorTranslatorFunc) Translate(ctx context.Context, err *RuleError) error { return f(ctx, err) }

type errorTranslatorImpl int

func (t errorTranslatorImpl) Translate(_ context.Context, err *RuleError) error { return err }

const DefaultErrorTranslator = errorTranslatorImpl(1)

var globalErrorTranslator ErrorTranslator = DefaultErrorTranslator
var globalErrorTranslatorLock sync.RWMutex

func SetErrorTranslator(translator ErrorTranslator) {
	globalErrorTranslatorLock.Lock()
	defer globalErrorTranslatorLock.Unlock()
	globalErrorTranslator = translator
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

func (k *KeyError) Error() string { return k.String() }

func (k *KeyError) String() string {
	b, err := k.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (k *KeyError) MarshalJSON() ([]byte, error) {
	aux := auxKeyError(*k)
	if _, ok := k.Err.(json.Marshaler); ok {
		k.Err = TextError(k.Err.Error())
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

func (e *Errors) Errs() []error { return e.errs }

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

// NopFunctionValidator does nothing and always returns nil. It's meant to be used as the first validator in a chain.
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

type BuilderFunc[T any] func(value T) Validator

func (f BuilderFunc[T]) Build(value T) Validator { return f(value) }

type FunctionValidator[T any] func(ctx context.Context, value T) error

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
		if err := WithTranslator[T](f)(ctx, value); err != nil {
			return err
		}
		return WithTranslator[T](g)(ctx, value)
	}
}

func WithTranslator[T any, Func FunctionValidatorConstraint[T]](f Func) Func {
	return func(ctx context.Context, value T) error {
		err := f(ctx, value)
		if err != nil {
			switch et := err.(type) {
			default:
				return err
			case *RuleError:
				return globalErrorTranslator.Translate(ctx, et)
			}
		}
		return nil
	}
}

func Named[T any, B Builder[T]](name string, value T, builder B) Validator {
	return ValidatorFunc(func(ctx context.Context) error {
		if err := builder.Build(value).Validate(ctx); err != nil {
			return NewKeyError(name, err)
		}
		return nil
	})
}

func Each[T any, V []T](fn func(each T) Validator) Builder[V] {
	return BuilderFunc[V](func(values V) Validator {
		return ValidatorFunc(func(ctx context.Context) error {
			errs := NewErrors()
			for _, each := range values {
				errs.Append(fn(each).Validate(ctx))
			}
			return errs.Err()
		})
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

type Pattern interface {
	RegExp() *regexp.Regexp
}
