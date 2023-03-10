package goval

import "context"

type MapValidator[K comparable, V any] func(ctx context.Context, values map[K]V) error

func Map[K comparable, V any]() MapValidator[K, V] {
	return NopValueValidator[map[K]V]
}

// Build attaches the value so the rules chain can consume it as an input that need to be validated.
func (mv MapValidator[K, V]) Build(values map[K]V) Validator {
	return validatorOf(mv, values)
}

// Required ensures the length is not 0 or the map is not nil.
func (mv MapValidator[K, V]) Required() MapValidator[K, V] {
	return Chain(mv, func(ctx context.Context, values map[K]V) error {
		if values == nil || len(values) == 0 {
			return Errorf("is required")
		}

		return nil
	})
}

// Min ensures the length is not less than the given min.
func (mv MapValidator[K, V]) Min(min int) MapValidator[K, V] {
	return Chain(mv, func(ctx context.Context, values map[K]V) error {
		if len(values) < min {
			return Errorf("length must be greater than %v", min)
		}
		return nil
	})
}

// Max ensures the length is not greater than the given max.
func (mv MapValidator[K, V]) Max(max int) MapValidator[K, V] {
	return Chain(mv, func(ctx context.Context, values map[K]V) error {
		if len(values) > max {
			return Errorf("length must be less than %v", max)
		}
		return nil
	})
}

func (mv MapValidator[K, V]) Each(builder Builder[V]) MapValidator[K, V] {
	return func(ctx context.Context, values map[K]V) error {
		for key, value := range values {
			if err := builder.Build(value).Validate(ctx); err != nil {
				return Errorf("key:%v, got: %s", key, err)
			}
		}
		return nil
	}
}
