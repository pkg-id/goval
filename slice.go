package goval

import "context"

// SliceValidator is a FunctionValidator that validates slices.
type SliceValidator[T any, V []T] FunctionValidator[V]

// Slice returns a SliceValidator with no rules.
// T is the type of the slice elements, V is the type of the slice.
func Slice[T any, V []T]() SliceValidator[T, V] {
	return NopFunctionValidator[V]()
}

// Build builds the validator chain and attaches the value to it.
func (sv SliceValidator[T, V]) Build(values V) Validator {
	return validatorOf(sv, values)
}

// With attaches the next rule to the chain.
func (sv SliceValidator[T, V]) With(next SliceValidator[T, V]) SliceValidator[T, V] {
	return Chain(sv, next)
}

// Required ensures the slice is not empty.
func (sv SliceValidator[T, V]) Required() SliceValidator[T, V] {
	return sv.With(func(ctx context.Context, values V) error {
		if len(values) == 0 {
			return NewRuleError(SliceRequired, values)
		}
		return nil
	})
}

// Min ensures the length of the slice is not less than the given min.
func (sv SliceValidator[T, V]) Min(min int) SliceValidator[T, V] {
	return sv.With(func(ctx context.Context, values V) error {
		if len(values) < min {
			return NewRuleError(SliceMin, values, min)
		}
		return nil
	})
}

// Max ensures the length of the slice is not greater than the given max.
func (sv SliceValidator[T, V]) Max(max int) SliceValidator[T, V] {
	return sv.With(func(ctx context.Context, values V) error {
		if len(values) > max {
			return NewRuleError(SliceMax, values, max)
		}
		return nil
	})
}

// Each ensures each element of the slice is satisfied by the given validator.
func (sv SliceValidator[T, V]) Each(validator Builder[T]) SliceValidator[T, V] {
	return func(ctx context.Context, values V) error {
		errs := make(Errors, 0)
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
		return errs.NilOrErr()
	}
}
