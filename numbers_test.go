package goval_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/pkg-id/goval"
)

func TestNumber(t *testing.T) {
	t.Run("int", NumberValidatorTestFunc(1, 0))
	t.Run("int8", NumberValidatorTestFunc[int8](1, 0))
	t.Run("int16", NumberValidatorTestFunc[int16](1, 0))
	t.Run("int32", NumberValidatorTestFunc[int32](1, 0))
	t.Run("int64", NumberValidatorTestFunc[int64](1, 0))

	t.Run("uint", NumberValidatorTestFunc[uint](1, 0))
	t.Run("uint8", NumberValidatorTestFunc[uint8](1, 0))
	t.Run("uint16", NumberValidatorTestFunc[uint16](1, 0))
	t.Run("uint32", NumberValidatorTestFunc[uint32](1, 0))
	t.Run("uint64", NumberValidatorTestFunc[uint64](1, 0))

	t.Run("float32", NumberValidatorTestFunc[float32](1.123, 0.0))
	t.Run("float64", NumberValidatorTestFunc(1.123, 0.0))
}

func NumberValidatorTestFunc[T goval.NumberConstraint](ok, fail T) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Number[T]().Validate(ctx, ok)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Validate(ctx, fail)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}
	}
}

func TestNilValidator_Required(t *testing.T) {
	t.Run("int", NumberValidatorRequiredTestFunc(1, 0))
	t.Run("int8", NumberValidatorRequiredTestFunc[int8](1, 0))
	t.Run("int16", NumberValidatorRequiredTestFunc[int16](1, 0))
	t.Run("int32", NumberValidatorRequiredTestFunc[int32](1, 0))
	t.Run("int64", NumberValidatorRequiredTestFunc[int64](1, 0))

	t.Run("uint", NumberValidatorRequiredTestFunc[uint](1, 0))
	t.Run("uint8", NumberValidatorRequiredTestFunc[uint8](1, 0))
	t.Run("uint16", NumberValidatorRequiredTestFunc[uint16](1, 0))
	t.Run("uint32", NumberValidatorRequiredTestFunc[uint32](1, 0))
	t.Run("uint64", NumberValidatorRequiredTestFunc[uint64](1, 0))

	t.Run("float32", NumberValidatorRequiredTestFunc[float32](1.123, 0.0))
	t.Run("float64", NumberValidatorRequiredTestFunc(1.123, 0.0))
}

func NumberValidatorRequiredTestFunc[T goval.NumberConstraint](ok, fail T) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Number[T]().Required().Validate(ctx, ok)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Required().Validate(ctx, fail)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberRequired) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberRequired, exp.Code)
		}

		if exp.Args != nil {
			t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
		}
	}
}

func TestNumberValidator_Min(t *testing.T) {
	t.Run("int", NumberValidatorMinTestFunc(3, 0))
	t.Run("int8", NumberValidatorMinTestFunc[int8](3, 0))
	t.Run("int16", NumberValidatorMinTestFunc[int16](3, 0))
	t.Run("int32", NumberValidatorMinTestFunc[int32](3, 0))
	t.Run("int64", NumberValidatorMinTestFunc[int64](3, 0))

	t.Run("uint", NumberValidatorMinTestFunc[uint](3, 0))
	t.Run("uint8", NumberValidatorMinTestFunc[uint8](3, 0))
	t.Run("uint16", NumberValidatorMinTestFunc[uint16](3, 0))
	t.Run("uint32", NumberValidatorMinTestFunc[uint32](3, 0))
	t.Run("uint64", NumberValidatorMinTestFunc[uint64](3, 0))

	t.Run("float32", NumberValidatorMinTestFunc[float32](3.123, 0.0))
	t.Run("float64", NumberValidatorMinTestFunc(3.123, 0.0))
}

func NumberValidatorMinTestFunc[T goval.NumberConstraint](ok, fail T) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Number[T]().Min(3).Validate(ctx, ok)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Min(3).Validate(ctx, fail)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberMin) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMin, exp.Code)
		}

		args := []any{T(3)}
		if !reflect.DeepEqual(exp.Args, args) {
			t.Errorf("expect the error args: %v, type: %T; got error args: %v, type: %T", args, args, exp.Args, exp.Args)
		}
	}
}

func TestNumberValidator_Max(t *testing.T) {
	t.Run("int", NumberValidatorMaxTestFunc(3, 5))
	t.Run("int8", NumberValidatorMaxTestFunc[int8](3, 5))
	t.Run("int16", NumberValidatorMaxTestFunc[int16](3, 5))
	t.Run("int32", NumberValidatorMaxTestFunc[int32](3, 5))
	t.Run("int64", NumberValidatorMaxTestFunc[int64](3, 5))

	t.Run("uint", NumberValidatorMaxTestFunc[uint](3, 5))
	t.Run("uint8", NumberValidatorMaxTestFunc[uint8](3, 5))
	t.Run("uint16", NumberValidatorMaxTestFunc[uint16](3, 5))
	t.Run("uint32", NumberValidatorMaxTestFunc[uint32](3, 5))
	t.Run("uint64", NumberValidatorMaxTestFunc[uint64](3, 5))

	t.Run("float32", NumberValidatorMaxTestFunc[float32](3.0, 3.1))
	t.Run("float64", NumberValidatorMaxTestFunc(3.0, 3.1))
}

func NumberValidatorMaxTestFunc[T goval.NumberConstraint](ok, fail T) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Number[T]().Max(3).Validate(ctx, ok)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Max(3).Validate(ctx, fail)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberMax) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMax, exp.Code)
		}

		args := []any{T(3)}
		if !reflect.DeepEqual(exp.Args, args) {
			t.Errorf("expect the error args: %v, type: %T; got error args: %v, type: %T", args, args, exp.Args, exp.Args)
		}
	}
}

func TestNumberValidator_In(t *testing.T) {
	t.Run("int", NumberValidatorInTestFunc(3, []int{1, 3}, []int{2, 4}))
	t.Run("int8", NumberValidatorInTestFunc(3, []int8{1, 3}, []int8{2, 4}))
	t.Run("int16", NumberValidatorInTestFunc(3, []int16{1, 3}, []int16{2, 4}))
	t.Run("int32", NumberValidatorInTestFunc(3, []int32{1, 3}, []int32{2, 4}))
	t.Run("int64", NumberValidatorInTestFunc(3, []int64{1, 3}, []int64{2, 4}))

	t.Run("uint", NumberValidatorInTestFunc(3, []uint{1, 3}, []uint{2, 4}))
	t.Run("uint8", NumberValidatorInTestFunc(3, []uint8{1, 3}, []uint8{2, 4}))
	t.Run("uint16", NumberValidatorInTestFunc(3, []uint16{1, 3}, []uint16{2, 4}))
	t.Run("uint32", NumberValidatorInTestFunc(3, []uint32{1, 3}, []uint32{2, 4}))
	t.Run("uint64", NumberValidatorInTestFunc(3, []uint64{1, 3}, []uint64{2, 4}))

	t.Run("float32", NumberValidatorInTestFunc(3.0, []float32{3.0, 2.0}, []float32{3.01}))
	t.Run("float64", NumberValidatorInTestFunc(3.0, []float64{3.0, 2.0}, []float64{3.01}))
}

func NumberValidatorInTestFunc[T goval.NumberConstraint, V []T](num T, ok V, fail V) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Number[T]().In(ok...).Validate(ctx, num)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().In(fail...).Validate(ctx, num)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberIn) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberIn, exp.Code)
		}

		args := []any{fail}
		if !reflect.DeepEqual(exp.Args, args) {
			t.Errorf("expect the error args: %v, type: %T; got error args: %v, type: %T", args, args, exp.Args, exp.Args)
		}
	}
}

func TestNumberValidator_When(t *testing.T) {
	isOdd := func(val int) bool { return val%2 == 1 }
	isEven := func(val int) bool { return val%2 == 0 }

	validator := goval.Number[int]().Required().Min(5).
		When(isEven, func(chain goval.NumberValidator[int]) goval.NumberValidator[int] { return chain.Min(12).Max(16) }).
		When(isOdd, func(chain goval.NumberValidator[int]) goval.NumberValidator[int] { return chain.Max(9) }).
		In(7, 16)

	t.Run("required", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, 0)

		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T, got type %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberRequired) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberRequired, exp.Code)
		}
	})

	t.Run("min at parent", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, 3)

		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T, got type %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberMin) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMin, exp.Code)
		}
	})

	t.Run("max when is odd", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, 11)

		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T, got type %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberMax) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMax, exp.Code)
		}
	})

	t.Run("min when is even", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, 10)

		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T, got type %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberMin) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMin, exp.Code)
		}
	})

	t.Run("max when is even", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, 20)

		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T, got type %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberMax) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMax, exp.Code)
		}
	})

	t.Run("in rules at parent", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, 14)

		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T, got type %T", exp, err)
		}

		if !exp.Code.Equal(goval.NumberIn) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberIn, exp.Code)
		}
	})

	t.Run("valid", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, 7)
		if err != nil {
			t.Fatalf("expect no error but got an error: %v", err)
		}
	})

}

func BenchmarkNumberValidator_Required(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 0
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Required().Validate(ctx, val)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 1
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Required().Validate(ctx, val)
		}
	})
}

func BenchmarkNumberValidator_Min(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Min(10).Validate(ctx, val)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Min(9).Validate(ctx, val)
		}
	})
}

func BenchmarkNumberValidator_Max(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Min(8).Validate(ctx, val)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Max(9).Validate(ctx, val)
		}
	})
}
