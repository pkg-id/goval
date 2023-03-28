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
		err := goval.Number[T]().Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Build(fail).Validate(ctx)
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
		err := goval.Number[T]().Required().Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Required().Build(fail).Validate(ctx)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !goval.IsCodeEqual(exp.Code, goval.NumberRequired) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberRequired, exp.Code)
		}

		inp, ok := exp.Input.(T)
		if !ok {
			t.Fatalf("expect the error input type: %T; got error input: %T", fail, exp.Input)
		}

		if inp != fail {
			t.Errorf("expect the error input value: %v; got error input value: %v", fail, inp)
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
		err := goval.Number[T]().Min(3).Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Min(3).Build(fail).Validate(ctx)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !goval.IsCodeEqual(exp.Code, goval.NumberMin) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMin, exp.Code)
		}

		inp, ok := exp.Input.(T)
		if !ok {
			t.Fatalf("expect the error input type: %T; got error input: %T", fail, exp.Input)
		}

		if inp != fail {
			t.Errorf("expect the error input value: %v; got error input value: %v", fail, inp)
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
		err := goval.Number[T]().Max(3).Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Number[T]().Max(3).Build(fail).Validate(ctx)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !goval.IsCodeEqual(exp.Code, goval.NumberMax) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.NumberMax, exp.Code)
		}

		inp, ok := exp.Input.(T)
		if !ok {
			t.Fatalf("expect the error input type: %T; got error input: %T", fail, exp.Input)
		}

		if inp != fail {
			t.Errorf("expect the error input value: %v; got error input value: %v", fail, inp)
		}

		args := []any{T(3)}
		if !reflect.DeepEqual(exp.Args, args) {
			t.Errorf("expect the error args: %v, type: %T; got error args: %v, type: %T", args, args, exp.Args, exp.Args)
		}
	}
}

func BenchmarkNumberValidator_Required(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 0
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Required().Build(val).Validate(ctx)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 1
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Required().Build(val).Validate(ctx)
		}
	})
}

func BenchmarkNumberValidator_Min(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Min(10).Build(val).Validate(ctx)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Min(9).Build(val).Validate(ctx)
		}
	})
}

func BenchmarkNumberValidator_Max(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Min(8).Build(val).Validate(ctx)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 9
		for i := 0; i < b.N; i++ {
			_ = goval.Number[int]().Max(9).Build(val).Validate(ctx)
		}
	})
}
