# Design Proposal

This document is a proposal for the design of the Go Validation Library (goval). It is used to discuss the design of the library and gather feedback from the team.

## What we want to achieve?

The goal of this library is to provide a simple and easy-to-use validation library for Go that is easy to integrate into existing projects and easy for developers to use. The library should be easy to extend and provide a simple way to add custom validators. No magic should happen behind the scenes. The library should be easy to understand and debug.

This library enhances the capabilities of the Go function as the first-class citizen and must be safe for concurrent use. The library is also designed to avoid using reflection as much as possible.

## Core Concepts

This package is designed to encourage the power of functional programming. Everything in this package is a function, and the functions are designed to be chained together to create complex validation logic. That means each validation should be defined as simply as possible and should be reusable/composable.

There are a few terms used in this library:

### Validator

Validator is an interface that defines the Validate function. The Validate function takes only `context.Context` as an argument and returns an error. The error can be nil if the validation is successful. **The input that needs to be validated must be provided by the implementation of the validator**. This validator is only used to group the validators together.


```go
type Validator interface {
	// Validate starts the validation process and returns an error if the validation fails.
    Validate(ctx context.Context) error
}
```

### Builder

Builder is an interface that defines the Validator from the ValueValidator. The Builder interface is used to construct the validator that can be used to start the validation process.

This Builder acts like a factory for Validator and also as the input supplier for the ValueValidator.

```go

type Builder[T any] interface {
    // Build returns a Validator that already has configured rules.
    Build(value T) Validator
}

```

### Chainer

Chainer is a function for chaining two ValueValidators into a single ValueValidator. The Chainer is used to compose the validation logic.

**This chaining process must be delayed until the invocation is triggered by the implementation of the Validator**. This ensures that the validation logic is not executed until it is needed and makes the validation logic reusable and composable in nature.

```go

// Chain creates a new function that chains the execution of two given functions into a single function.
// Here's an example: suppose we have two functions:
//
//	var f func(ctx context.Context, value T) error
//	var g func(ctx context.Context, value T) error
//
// We want to combine `f` and `g` into a single function, but without executing them immediately. In other words, we
// want to delay the execution of `f` and `g` until the new function is executed.
// Let's call the new function `h`. When `h` is executed, `f` will be executed first. If `f` executes without error,
// then `g` will be executed next. If `h` returns any error it will be an error that returned either `f` or `g`.
func Chain[T any](f, g func(ctx context.Context, value T) error) func(ctx context.Context, value T) error {
	return func(ctx context.Context, value T) error {
		if err := f(ctx, value); err != nil {
			return err
		}

		return g(ctx, value)
	}
}

```

### ValueValidator

`ValueValidator` is the actual validator that validates the input. But this implementation is not able to start the validation process. Instead, this `ValueValidator` implement the `Builder`.

This `ValueValidator` is used to compose the validation logic by chaining a series of validation rules together. The validation rules are also defined in this implementation. Validation rules are basically a function (or a method of the `ValueValidator`) that returns a `ValueValidator` and composes the previous validation rules with the new validation rules.

```go

// StringValidator is a ValueValidator that validates the string value.
// Each ValueValidator must follow this function signature.
//   - type ValueValidator[T any] func(ctx context.Context, value T) error
type StringValidator func(ctx context.Context, value string) error

// String is the constructor for the StringValidator.
func String(value string) StringValidator {
    // This returned function is not validating anything. It just returns nil. This called as No-Op function.
    // We will discuss this later.
    return func(ctx context.Context, value string) error {
        return nil
    }
}

// Build implements the Builder interface.
func (f StringValidator) Build(value string) Validator {
	// This function is an adapter function that converts the ValueValidator to the Validator.
    return ValidatorFunc(func(ctx context.Context) error {
        return f(ctx, value)
    })
}

func (f StringValidator) Required() StringValidator {
	// This is where composition happens. It chains the current validation rule `f` with the new validation rule.
    return Chain(f, func(ctx context.Context, value string) error {
        // This is the actual validation logic.
        // This function is validating the string value.
        // If the value is empty, then it returns an error.
        if value == "" {
            return errors.New("value is empty")
        }

		// When the validation is successful, it returns nil.
        return nil
    })
}

func (f StringValidator) Min(min int) StringValidator {
	// Same as what we did in the `Required` function.
    return Chain(f, func(ctx context.Context, value string) error {
        if len(value) < min {
            return fmt.Errorf("value length is less than %d", min)
        }
        return nil
    })
}

```

As we can see, there are one function and three methods are defined in the code above.

1. The `String` function or the constructor for the `StringValidator`.
2. The `Build` method that implements the `Builder` interface.
3. The `Required` and `Min` are the validation rules that are used to compose the validation logic.

Those patterns will be look similar for other types of validators.

If you remember, previously we said that constructor created a `ValueValidator` that does nothing. This is the reason why we need to create a `ValueValidator` that does nothing. This is because we want to compose the validation logic by chaining a series of validation rules together. But, the `Chain` function is required at least two function. So, by creating a `ValueValidator` that does nothing, we can use it as the first argument for the `Chain` function and ensure that the `Chain` function will always have at least two functions when it chained with available validation rules. And, this also to ensure that a `ValueValidator` without any validation rules will always return nil since no rules are violated.