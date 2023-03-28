package goval

import "context"

// MapValidator is a FunctionValidator that validates map[K]V.
type MapValidator[K comparable, V any] FunctionValidator[map[K]V]

// Map returns a MapValidator with no rules.
func Map[K comparable, V any]() MapValidator[K, V] {
	return NopFunctionValidator[map[K]V]()
}

// Build builds the validator chain and attaches the value to it.
func (mv MapValidator[K, V]) Build(values map[K]V) Validator {
	return validatorOf(mv, values)
}

// With attaches the next rule to the chain.
func (mv MapValidator[K, V]) With(next MapValidator[K, V]) MapValidator[K, V] {
	return Chain(mv, next)
}

// Required ensures the length is not zero.
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
			return NewRuleError(MapMin, values, min)
		}
		return nil
	})
}

// Max ensures the length is not greater than the given max.
func (mv MapValidator[K, V]) Max(max int) MapValidator[K, V] {
	return mv.With(func(ctx context.Context, values map[K]V) error {
		if len(values) > max {
			return NewRuleError(MapMax, values, max)
		}
		return nil
	})
}

// Each ensures each element of the map is satisfied by the given validator.
func (mv MapValidator[K, V]) Each(validator Builder[V]) MapValidator[K, V] {
	return func(ctx context.Context, values map[K]V) error {
		errs := make([]error, 0)
		for _, value := range values {
			if err := validator.Build(value).Validate(ctx); err != nil {
				switch et := err.(type) {
				default:
					return err
				case *RuleError:
					errs = append(errs, et)
				}
			}
		}

		if len(errs) > 0 {
			return NewErrors(errs)
		}
		return nil
	}
}
