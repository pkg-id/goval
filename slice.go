package goval

import "context"

type SliceValidator[T any, V []T] func(ctx context.Context, values V) error

func Slice[T any, V []T]() SliceValidator[T, V] {
	return NopValueValidator[V]
}

// Build attaches the value so the rules chain can consume it as an input that need to be validated.
func (nv SliceValidator[T, V]) Build(values V) Validator {
	return validatorOf(nv, values)
}

// Required ensures the number is not a zero value.
func (nv SliceValidator[T, V]) Required() SliceValidator[T, V] {
	return Chain(nv, func(ctx context.Context, values V) error {
		if values == nil || len(values) == 0 {
			return Errorf("is required")
		}

		return nil
	})
}

// Min ensures the number is not less than the given min.
func (nv SliceValidator[T, V]) Min(min int) SliceValidator[T, V] {
	return Chain(nv, func(ctx context.Context, values V) error {
		if len(values) < min {
			return Errorf("length must be greater than %v", min)
		}
		return nil
	})
}

// Max ensures the number is not greater than the given max.
func (nv SliceValidator[T, V]) Max(max int) SliceValidator[T, V] {
	return Chain(nv, func(ctx context.Context, values V) error {
		if len(values) > max {
			return Errorf("length must be less than %v", max)
		}
		return nil
	})
}

func (nv SliceValidator[T, V]) Each(builder Builder[T]) SliceValidator[T, V] {
	return func(ctx context.Context, values V) error {
		for index, value := range values {
			if err := builder.Build(value).Validate(ctx); err != nil {
				return Errorf("index:%d, got: %s", index, err)
			}
		}
		return nil
	}
}
