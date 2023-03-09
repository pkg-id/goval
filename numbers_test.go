package goval_test

import (
	"context"
	"github.com/pkg-id/goval"
	"testing"
)

func TestNumberValidator_Required(t *testing.T) {
	t.Run("required_when_error", func(t *testing.T) {
		constraintNumberRequiredWhenError[int](t, "int")
		constraintNumberRequiredWhenError[int8](t, "int8")
		constraintNumberRequiredWhenError[int16](t, "int16")
		constraintNumberRequiredWhenError[int32](t, "int32")
		constraintNumberRequiredWhenError[int64](t, "int64")
		constraintNumberRequiredWhenError[uint](t, "uint")
		constraintNumberRequiredWhenError[uint8](t, "unit8")
		constraintNumberRequiredWhenError[uint16](t, "uint16")
		constraintNumberRequiredWhenError[uint32](t, "uint32")
		constraintNumberRequiredWhenError[uint64](t, "uint64")
		constraintNumberRequiredWhenError[float32](t, "float32")
		constraintNumberRequiredWhenError[float64](t, "float64")
	})

	t.Run("required_when_ok", func(t *testing.T) {
		constraintNumberRequiredWhenOK[int](t, "int")
		constraintNumberRequiredWhenOK[int8](t, "int8")
		constraintNumberRequiredWhenOK[int16](t, "int16")
		constraintNumberRequiredWhenOK[int32](t, "int32")
		constraintNumberRequiredWhenOK[int64](t, "int64")
		constraintNumberRequiredWhenOK[uint](t, "uint")
		constraintNumberRequiredWhenOK[uint8](t, "unit8")
		constraintNumberRequiredWhenOK[uint16](t, "uint16")
		constraintNumberRequiredWhenOK[uint32](t, "uint32")
		constraintNumberRequiredWhenOK[uint64](t, "uint64")
		constraintNumberRequiredWhenOK[float32](t, "float32")
		constraintNumberRequiredWhenOK[float64](t, "float64")
	})
}

func TestNumberValidator_Min(t *testing.T) {
	t.Run("min_when_error", func(t *testing.T) {
		constraintNumberMinWhenError[int](t, "int")
		constraintNumberMinWhenError[int8](t, "int8")
		constraintNumberMinWhenError[int16](t, "int16")
		constraintNumberMinWhenError[int32](t, "int32")
		constraintNumberMinWhenError[int64](t, "int64")
		constraintNumberMinWhenError[uint](t, "uint")
		constraintNumberMinWhenError[uint8](t, "unit8")
		constraintNumberMinWhenError[uint16](t, "uint16")
		constraintNumberMinWhenError[uint32](t, "uint32")
		constraintNumberMinWhenError[uint64](t, "uint64")
		constraintNumberMinWhenError[float32](t, "float32")
		constraintNumberMinWhenError[float64](t, "float64")
	})

	t.Run("min_when_ok", func(t *testing.T) {
		constraintNumberMinWhenOK[int](t, "int")
		constraintNumberMinWhenOK[int8](t, "int8")
		constraintNumberMinWhenOK[int16](t, "int16")
		constraintNumberMinWhenOK[int32](t, "int32")
		constraintNumberMinWhenOK[int64](t, "int64")
		constraintNumberMinWhenOK[uint](t, "uint")
		constraintNumberMinWhenOK[uint8](t, "unit8")
		constraintNumberMinWhenOK[uint16](t, "uint16")
		constraintNumberMinWhenOK[uint32](t, "uint32")
		constraintNumberMinWhenOK[uint64](t, "uint64")
		constraintNumberMinWhenOK[float32](t, "float32")
		constraintNumberMinWhenOK[float64](t, "float64")
	})
}

func TestNumberValidator_Max(t *testing.T) {
	t.Run("max_when_error", func(t *testing.T) {
		constraintNumberMaxWhenError[int](t, "int")
		constraintNumberMaxWhenError[int8](t, "int8")
		constraintNumberMaxWhenError[int16](t, "int16")
		constraintNumberMaxWhenError[int32](t, "int32")
		constraintNumberMaxWhenError[int64](t, "int64")
		constraintNumberMaxWhenError[uint](t, "uint")
		constraintNumberMaxWhenError[uint8](t, "unit8")
		constraintNumberMaxWhenError[uint16](t, "uint16")
		constraintNumberMaxWhenError[uint32](t, "uint32")
		constraintNumberMaxWhenError[uint64](t, "uint64")
		constraintNumberMaxWhenError[float32](t, "float32")
		constraintNumberMaxWhenError[float64](t, "float64")
	})

	t.Run("max_when_ok", func(t *testing.T) {
		constraintNumberMaxWhenOK[int](t, "int")
		constraintNumberMaxWhenOK[int8](t, "int8")
		constraintNumberMaxWhenOK[int16](t, "int16")
		constraintNumberMaxWhenOK[int32](t, "int32")
		constraintNumberMaxWhenOK[int64](t, "int64")
		constraintNumberMaxWhenOK[uint](t, "uint")
		constraintNumberMaxWhenOK[uint8](t, "unit8")
		constraintNumberMaxWhenOK[uint16](t, "uint16")
		constraintNumberMaxWhenOK[uint32](t, "uint32")
		constraintNumberMaxWhenOK[uint64](t, "uint64")
		constraintNumberMaxWhenOK[float32](t, "float32")
		constraintNumberMaxWhenOK[float64](t, "float64")
	})
}

func constraintNumberRequiredWhenError[T goval.NumberConstraint](t *testing.T, describe string) {
	t.Run(describe, func(t *testing.T) {
		var n T
		err := goval.Number[T]().Required().WithValue(n).Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}
	})
}

func constraintNumberRequiredWhenOK[T goval.NumberConstraint](t *testing.T, describe string) {
	t.Run(describe, func(t *testing.T) {
		var n T = 1
		err := goval.Number[T]().Required().WithValue(n).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func constraintNumberMinWhenError[T goval.NumberConstraint](t *testing.T, describe string) {
	t.Run(describe, func(t *testing.T) {
		var n T = 9
		err := goval.Number[T]().Required().Min(10).WithValue(n).Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}
	})
}

func constraintNumberMinWhenOK[T goval.NumberConstraint](t *testing.T, describe string) {
	t.Run(describe, func(t *testing.T) {
		var n T = 10
		err := goval.Number[T]().Required().Min(10).WithValue(n).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func constraintNumberMaxWhenError[T goval.NumberConstraint](t *testing.T, describe string) {
	t.Run(describe, func(t *testing.T) {
		var n T = 11
		err := goval.Number[T]().Required().Min(1).Max(10).WithValue(n).Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}
	})
}

func constraintNumberMaxWhenOK[T goval.NumberConstraint](t *testing.T, describe string) {
	t.Run(describe, func(t *testing.T) {
		var n T = 10
		err := goval.Number[T]().Required().Min(1).Max(10).WithValue(n).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}
