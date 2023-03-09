package goval

import "context"

// StringValidator is a validator for string type.
type StringValidator func(ctx context.Context, value string) error

// String is StringValidator constructor. This function is used to initialize
// the rules chain. Since, it will be a first rule in the chain, it not validates anything.
func String() StringValidator {
	return NopValueValidator[string]
}

// WithValue attaches the value so the rules chain can consume it as an input that need to be validated.
func (sv StringValidator) WithValue(value string) Validator {
	return validatorOf(sv, value)
}

// Required ensures the string is not an empty string.
func (sv StringValidator) Required() StringValidator {
	return Chain(sv, func(ctx context.Context, value string) error {
		if value == "" {
			return Error("is required")
		}
		return nil
	})
}

// Min ensures the length of the string is not less than the given length.
func (sv StringValidator) Min(length int) StringValidator {
	return Chain(sv, func(ctx context.Context, value string) error {
		if len(value) < length {
			return Errorf("length must be at least %d characters", length)
		}
		return nil
	})
}

// Max ensures the length of the string is not greater than the given length.
func (sv StringValidator) Max(length int) StringValidator {
	return Chain(sv, func(ctx context.Context, value string) error {
		if len(value) > length {
			return Errorf("length must be less than %d characters", length)
		}
		return nil
	})
}
