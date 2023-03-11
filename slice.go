package goval

import "context"

type SliceValidator[T any, V []T] func(ctx context.Context, values V) error

func Slice[T any, V []T]() SliceValidator[T, V] {
	return NopValueValidator[V]
}

// Build attaches the value so the rules chain can consume it as an input that need to be validated.
func (sv SliceValidator[T, V]) Build(values V) Validator {
	return validatorOf(sv, values)
}

// Required ensures the number is not a zero value.
func (sv SliceValidator[T, V]) Required() SliceValidator[T, V] {
	return Chain(sv, func(ctx context.Context, values V) error {
		if values == nil || len(values) == 0 {
			return NewRuleError(SliceRequired, values)
		}
		return nil
	})
}

// Min ensures the number is not less than the given min.
func (sv SliceValidator[T, V]) Min(min int) SliceValidator[T, V] {
	return Chain(sv, func(ctx context.Context, values V) error {
		if len(values) < min {
			return NewRuleError(SliceMin, values, min)
		}
		return nil
	})
}

// Max ensures the number is not greater than the given max.
func (sv SliceValidator[T, V]) Max(max int) SliceValidator[T, V] {
	return Chain(sv, func(ctx context.Context, values V) error {
		if len(values) > max {
			return NewRuleError(SliceMax, values, max)
		}
		return nil
	})
}

func (sv SliceValidator[T, V]) Each(builder Builder[T]) SliceValidator[T, V] {
	return func(ctx context.Context, values V) error {
		errs := make([]RuleError, 0)
		for _, value := range values {
			if err := builder.Build(value).Validate(ctx); err != nil {
				switch et := err.(type) {
				default:
					return err
				case *RuleError:
					errs = append(errs, *et)
				}
			}
		}

		if len(errs) > 0 {
			return NewRuleErrors(SliceEach, errs, values)
		}
		return nil
	}
}
