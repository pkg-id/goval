package goval

import (
	"context"
	"encoding/json"
	"sync"
)

func NewRuleError(code RuleCode, input any, args ...any) *RuleError {
	return &RuleError{
		Code:  code,
		Input: input,
		Args:  args,
	}
}

type auxRuleError *RuleError

type RuleError struct {
	Code  RuleCode `json:"code"`
	Input any      `json:"input"`
	Args  []any    `json:"args,omitempty"`
}

func (r *RuleError) Error() string                { return r.String() }
func (r *RuleError) String() string               { return stringifyJSON(r) }
func (r *RuleError) MarshalJSON() ([]byte, error) { return json.Marshal(auxRuleError(r)) }

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

func (e *RuleErrors) Error() string                { return e.String() }
func (e *RuleErrors) String() string               { return stringifyJSON(e) }
func (e *RuleErrors) MarshalJSON() ([]byte, error) { return json.Marshal(e.Errs) }

// TextError is an error type for turning an ordinary string to an error.
type TextError string

func (t TextError) Error() string                { return t.String() }
func (t TextError) String() string               { return string(t) }
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

type auxKeyError *KeyError

func NewKeyError(key string, err error) *KeyError {
	return &KeyError{
		Key: key,
		Err: err,
	}
}

func (k *KeyError) Error() string  { return k.String() }
func (k *KeyError) String() string { return stringifyJSON(k) }

func (k *KeyError) MarshalJSON() ([]byte, error) {
	aux := auxKeyError(k)
	if _, ok := k.Err.(json.Marshaler); !ok {
		k.Err = TextError(k.Err.Error())
	}
	return json.Marshal(aux)
}

// Errors is a type for collecting multiple errors and bundling them into a single error.
type Errors struct {
	errs []error
}

// NewErrors creates a new error collector.
func NewErrors(errs []error) *Errors {
	return &Errors{errs: errs}
}

func (e *Errors) Error() string                { return e.String() }
func (e *Errors) String() string               { return stringifyJSON(e) }
func (e *Errors) MarshalJSON() ([]byte, error) { return json.Marshal(e.errs) }

func stringifyJSON(m json.Marshaler) string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}
