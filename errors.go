package goval

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// jsonErrorStringer is an interface that combines error, json.Marshaler and fmt.Stringer.
type jsonErrorStringer interface {
	error
	json.Marshaler
	fmt.Stringer
}

// auxRuleError is an auxiliary type for marshaling RuleError.
// It is used to avoid infinite recursion when marshaling RuleError.
type auxRuleError RuleError

// RuleError is an error type for validation errors.
type RuleError struct {
	Code  RuleCoder `json:"code"`           // the error code that identifies which rule failed.
	Input any       `json:"input"`          // the actual value that failed the validation.
	Args  []any     `json:"args,omitempty"` // additional arguments for the error.
}

// ensure RuleError implements jsonErrorStringer.
var _ jsonErrorStringer = (*RuleError)(nil)

// NewRuleError creates a new RuleError.
func NewRuleError(code RuleCoder, input any, args ...any) *RuleError {
	return &RuleError{
		Code:  code,
		Input: input,
		Args:  args,
	}
}

func (r *RuleError) Error() string                { return r.String() }
func (r *RuleError) String() string               { return stringifyJSON(r) }
func (r *RuleError) MarshalJSON() ([]byte, error) { return json.Marshal(auxRuleError(*r)) }

// TextError is an error type for turning an ordinary string to an error.
// This error type is intended to be used for creating an error that can be marshaled to JSON.
// For example, when overriding th ErrorTranslator, the implementation requires to return an error,
// you can use TextError to create an error message from a string literal.
type TextError string

var _ jsonErrorStringer = TextError("")

func (t TextError) Error() string                { return t.String() }
func (t TextError) String() string               { return string(t) }
func (t TextError) MarshalJSON() ([]byte, error) { return json.Marshal(t.String()) }

// ErrorTranslator is an interface for translating RuleError to a readable error.
type ErrorTranslator interface {
	// Translate translates a RuleError to a readable error.
	Translate(ctx context.Context, err *RuleError) error
}

// errorTranslatorImpl is an implementation of ErrorTranslator.
type errorTranslatorImpl int

// Translate implements ErrorTranslator.
func (t errorTranslatorImpl) Translate(_ context.Context, err *RuleError) error { return err }

// DefaultErrorTranslator is the default ErrorTranslator that never translates the error.
// In other words, it always returns the original error.
const DefaultErrorTranslator = errorTranslatorImpl(1)

var globalErrorTranslator ErrorTranslator = DefaultErrorTranslator
var globalErrorTranslatorLock sync.RWMutex

// SetErrorTranslator sets the global ErrorTranslator.
func SetErrorTranslator(translator ErrorTranslator) {
	globalErrorTranslatorLock.Lock()
	defer globalErrorTranslatorLock.Unlock()
	globalErrorTranslator = translator
}

// KeyError is an error with a key to give more context to the error.
type KeyError struct {
	Key string `json:"key"`
	Err error  `json:"err"`
}

// auxKeyError is an auxiliary type for marshaling KeyError.
type auxKeyError KeyError

// ensure KeyError implements jsonErrorStringer.
var _ jsonErrorStringer = (*KeyError)(nil)

// NewKeyError creates a new KeyError.
func NewKeyError(key string, err error) *KeyError {
	return &KeyError{
		Key: key,
		Err: err,
	}
}

func (k *KeyError) Error() string  { return k.String() }
func (k *KeyError) String() string { return stringifyJSON(k) }
func (k *KeyError) MarshalJSON() ([]byte, error) {
	// if the error is not a json.Marshaler, we convert it to a TextError.
	if _, ok := k.Err.(json.Marshaler); !ok {
		k.Err = TextError(k.Err.Error())
	}

	aux := auxKeyError(*k)
	return json.Marshal(aux)
}

// Errors is a type for collecting multiple errors and bundling them into a single error.
type Errors []error

// ensure Errors implements jsonErrorStringer.
var _ jsonErrorStringer = make(Errors, 0)

func (e Errors) Error() string                { return e.String() }
func (e Errors) String() string               { return stringifyJSON(e) }
func (e Errors) MarshalJSON() ([]byte, error) { return json.Marshal([]error(e)) }
func (e Errors) NilIfEmpty() error {
	if len(e) > 0 {
		return e
	}
	return nil
}

// stringifyJSON converts a json.Marshaler to a string.
// If the json.Marshaler returns an error, the error is returned as a string.
func stringifyJSON(m json.Marshaler) string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}
