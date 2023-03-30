package goval

import "context"

// PtrValidator is a FunctionValidator that validates *T.
type PtrValidator[T any] FunctionValidator[*T]

// Ptr returns a PtrValidator with no rules.
func Ptr[T any]() PtrValidator[T] { return NopFunctionValidator[*T]() }

// Build builds the validator chain and attaches the value to it.
func (f PtrValidator[T]) Build(value *T) Validator {
	return validatorOf(f, value)
}

// With attaches the next rule to the chain.
func (f PtrValidator[T]) With(next PtrValidator[T]) PtrValidator[T] {
	return Chain(f, next)
}

// Required ensures the pointer is not nil.
func (f PtrValidator[T]) Required() PtrValidator[T] {
	return f.With(func(ctx context.Context, value *T) error {
		if value == nil {
			return NewRuleError(PtrRequired)
		}
		return nil
	})
}

// Optional uses the given validator to validate the value if it is not nil.
func (f PtrValidator[T]) Optional(validator Builder[T]) PtrValidator[T] {
	return f.With(func(ctx context.Context, value *T) error {
		if value != nil {
			return f.Then(validator)(ctx, value)
		}
		return nil
	})
}

// Then chains the given validator to the current validator.
// It will be panic if the value of T is nil.
// Use Optional to optionally validate the value.
func (f PtrValidator[T]) Then(validator Builder[T]) PtrValidator[T] {
	return Chain(f, func(ctx context.Context, value *T) error {
		return validator.Build(*value).Validate(ctx)
	})
}
