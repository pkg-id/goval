package goval

import "context"

type MapValidator[K comparable, V any] FunctionValidator[map[K]V]

func Map[K comparable, V any]() MapValidator[K, V] {
	return NopFunctionValidator[map[K]V]()
}

// Build attaches the value so the rules chain can consume it as an input that need to be validated.
func (mv MapValidator[K, V]) Build(values map[K]V) Validator {
	return validatorOf(mv, values)
}

func (mv MapValidator[K, V]) With(next MapValidator[K, V]) MapValidator[K, V] {
	return Chain(mv, next)
}

// Required ensures the length is not 0 or the map is not nil.
func (mv MapValidator[K, V]) Required() MapValidator[K, V] {
	return mv.With(func(ctx context.Context, values map[K]V) error {
		if len(values) == 0 {
			return NewRuleError(MapRequired, values)
		}

		return nil
	})
}

// Min ensures the length is not less than the given min.
func (mv MapValidator[K, V]) Min(min int) MapValidator[K, V] {
	return mv.With(func(ctx context.Context, values map[K]V) error {
		if len(values) < min {
			return NewRuleError(MapMin, values)
		}
		return nil
	})
}

// Max ensures the length is not greater than the given max.
func (mv MapValidator[K, V]) Max(max int) MapValidator[K, V] {
	return mv.With(func(ctx context.Context, values map[K]V) error {
		if len(values) > max {
			return NewRuleError(MapMax, values)
		}
		return nil
	})
}

func (mv MapValidator[K, V]) Each(builder Builder[V]) MapValidator[K, V] {
	return func(ctx context.Context, values map[K]V) error {
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
			return NewRuleErrors(MapEach, errs, values)
		}
		return nil
	}
}
