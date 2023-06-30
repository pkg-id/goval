package goval

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg-id/goval/funcs"
)

// StringValidator is a FunctionValidator that validates string.
// For backward compatibility.
type StringValidator = StringVariantValidator[string]

// String returns a StringValidator with no rules.
// For backward compatibility.
func String() StringValidator {
	return StringVariant[string]()
}

// StringVariantValidator is a FunctionValidator that validates string variants.
type StringVariantValidator[T ~string] FunctionValidator[T]

// StringVariant returns a StringVariantValidator with no rules.
func StringVariant[T ~string]() StringVariantValidator[T] {
	return NopFunctionValidator[T]
}

// Validate executes the validation rules immediately.
func (f StringVariantValidator[T]) Validate(ctx context.Context, value T) error {
	return validatorOf(f, value).Validate(ctx)
}

// With attaches the next rule to the chain.
func (f StringVariantValidator[T]) With(next StringVariantValidator[T]) StringVariantValidator[T] {
	return Chain(f, next)
}

// Required ensures the string is not empty.
func (f StringVariantValidator[T]) Required() StringVariantValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		if value == "" {
			return NewRuleError(StringRequired)
		}
		return nil
	})
}

// Min ensures the length of the string is not less than the given length.
func (f StringVariantValidator[T]) Min(length int) StringVariantValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		if len(value) < length {
			return NewRuleError(StringMin, length)
		}
		return nil
	})
}

// Max ensures the length of the string is not greater than the given length.
func (f StringVariantValidator[T]) Max(length int) StringVariantValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		if len(value) > length {
			return NewRuleError(StringMax, length)
		}
		return nil
	})
}

// Match ensures the string matches the given pattern.
// If pattern cause panic, will be recovered.
func (f StringVariantValidator[T]) Match(pattern Pattern) StringVariantValidator[T] {
	return f.With(func(ctx context.Context, value T) (err error) {
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("panic: %v", rec)
			}
		}()

		exp := pattern.RegExp()
		if !exp.MatchString(string(value)) {
			return NewRuleError(StringMatch, exp.String())
		}

		return err
	})
}

// In ensures that the provided string is one of the specified options.
// This validation is case-sensitive, use InFold to perform a case-insensitive In validation.
func (f StringVariantValidator[T]) In(options ...T) StringVariantValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		ok := funcs.Contains(options, func(opt T) bool { return value == opt })
		if !ok {
			return NewRuleError(StringIn, options)
		}
		return nil
	})
}

// InFold ensures that the provided string is one of the specified options with case-insensitivity.
func (f StringVariantValidator[T]) InFold(options ...T) StringVariantValidator[T] {
	return f.With(func(ctx context.Context, value T) error {
		ok := funcs.Contains(options, func(opt T) bool { return strings.EqualFold(string(value), string(opt)) })
		if !ok {
			return NewRuleError(StringInFold, options)
		}
		return nil
	})
}

// When adds validation logic to the chain based on a condition for string values.
//
// If the predicate returns true, the result of the mapper function is added to the chain,
// and the input value is validated using the new chain. Otherwise, the original chain is returned unmodified.
//
// The mapper function takes a StringValidator instance and returns a new StringValidator instance with
// additional validation logic.
func (f StringVariantValidator[T]) When(p Predicate[T], m Mapper[T, StringVariantValidator[T]]) StringVariantValidator[T] {
	return whenLinker(f, p, m)
}
