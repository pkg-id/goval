package goval

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

type RuleError struct {
	Code  RuleCode `json:"code"`
	Input any      `json:"input"`
	Args  []any    `json:"args,omitempty"`
}

type auxRuleError RuleError

func NewRuleError(code RuleCode, input any, args ...any) *RuleError {
	return &RuleError{
		Code:  code,
		Input: input,
		Args:  args,
	}
}

func (r *RuleError) Error() string { return r.String() }

func (r *RuleError) String() string {
	b, err := r.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (r *RuleError) MarshalJSON() ([]byte, error) {
	aux := auxRuleError(*r)
	return json.Marshal(aux)
}

type RuleErrors struct {
	Code RuleCode
	Errs []RuleError
	Args []any
}

func NewRuleErrors(code RuleCode, errs []RuleError, args ...any) *RuleErrors {
	return &RuleErrors{
		Code: code,
		Errs: errs,
		Args: args,
	}
}

func (e *RuleErrors) Error() string { return e.String() }

func (e *RuleErrors) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		return "goval: RuleErrors.String: " + err.Error()
	}
	return string(b)
}

func (e *RuleErrors) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(e.Errs)
	if err != nil {
		return nil, fmt.Errorf("goval: RuleErrors.MarshalJSON: %w", err)
	}

	return b, nil
}

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
	if _, ok := k.Err.(json.Marshaler); !ok {
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
