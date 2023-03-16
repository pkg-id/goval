package goval

import "context"

type PtrValidator[T any] FunctionValidator[*T]

func Ptr[T any]() PtrValidator[T] { return NopFunctionValidator[*T]() }

func (f PtrValidator[T]) Build(value *T) Validator {
	return validatorOf(f, value)
}

func (f PtrValidator[T]) With(next PtrValidator[T]) PtrValidator[T] {
	return Chain(f, next)
}

func (f PtrValidator[T]) Required() PtrValidator[T] {
	return f.With(func(ctx context.Context, value *T) error {
		if value == nil {
			return NewRuleError(NilRequired, value)
		}
		return nil
	})
}

func (f PtrValidator[T]) Optional(builder Builder[T]) PtrValidator[T] {
	return f.With(func(ctx context.Context, value *T) error {
		if value != nil {
			return builder.Build(*value).Validate(ctx)
		}
		return nil
	})
}

func (f PtrValidator[T]) Next(builder Builder[T]) PtrValidator[T] {
	return Chain(f, func(ctx context.Context, value *T) error {
		return builder.Build(*value).Validate(ctx)
	})
}
