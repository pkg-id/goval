package goval_test

import (
	"context"
	"errors"
	"github.com/pkg-id/goval/govalregex"
	"reflect"
	"testing"

	"github.com/pkg-id/goval"
)

func TestSlice(t *testing.T) {
	t.Run("int", SliceValidatorTestFunc([]int{1}, []int{}))
	t.Run("int8", SliceValidatorTestFunc([]int8{1}, nil))
	t.Run("int16", SliceValidatorTestFunc([]int16{1}, []int16{}))
	t.Run("int32", SliceValidatorTestFunc([]int32{1}, nil))
	t.Run("int64", SliceValidatorTestFunc([]int64{1}, nil))

	t.Run("uint", SliceValidatorTestFunc([]uint{1}, nil))
	t.Run("uint8", SliceValidatorTestFunc([]uint8{1}, []uint8{}))
	t.Run("uint16", SliceValidatorTestFunc([]uint16{1}, nil))
	t.Run("uint32", SliceValidatorTestFunc([]uint32{1}, []uint32{}))
	t.Run("uint64", SliceValidatorTestFunc([]uint64{1}, nil))

	t.Run("float32", SliceValidatorTestFunc([]float32{1.123}, nil))
	t.Run("float64", SliceValidatorTestFunc([]float64{1.123}, []float64{}))

	t.Run("string", SliceValidatorTestFunc([]string{"abc"}, []string{}))

	t.Run("byte", SliceValidatorTestFunc([]byte{0x1}, []byte{}))

	t.Run("slice", SliceValidatorTestFunc([][]int{{1}}, [][]int{}))

	t.Run("map", SliceValidatorTestFunc([]map[int]int{{1: 1}}, []map[int]int{}))

	type X struct {
		A int
		B string
	}
	t.Run("struct", SliceValidatorRequiredTestFunc([]X{{A: 0, B: ""}}, []X{}))
}

func SliceValidatorTestFunc[T any, Slice []T](ok, fail Slice) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Slice[T]().Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Slice[T]().Build(fail).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}
	}
}

func TestSliceValidator_Required(t *testing.T) {
	t.Run("int", SliceValidatorRequiredTestFunc([]int{1}, []int{}))
	t.Run("int8", SliceValidatorRequiredTestFunc([]int8{1}, nil))
	t.Run("int16", SliceValidatorRequiredTestFunc([]int16{1}, []int16{}))
	t.Run("int32", SliceValidatorRequiredTestFunc([]int32{1}, nil))
	t.Run("int64", SliceValidatorRequiredTestFunc([]int64{1}, nil))

	t.Run("uint", SliceValidatorRequiredTestFunc([]uint{1}, nil))
	t.Run("uint8", SliceValidatorRequiredTestFunc([]uint8{1}, []uint8{}))
	t.Run("uint16", SliceValidatorRequiredTestFunc([]uint16{1}, nil))
	t.Run("uint32", SliceValidatorRequiredTestFunc([]uint32{1}, []uint32{}))
	t.Run("uint64", SliceValidatorRequiredTestFunc([]uint64{1}, nil))

	t.Run("float32", SliceValidatorRequiredTestFunc([]float32{1.123}, nil))
	t.Run("float64", SliceValidatorRequiredTestFunc([]float64{1.123}, []float64{}))

	t.Run("string", SliceValidatorRequiredTestFunc([]string{"abc"}, []string{}))

	t.Run("byte", SliceValidatorRequiredTestFunc([]byte{0x1}, []byte{}))

	t.Run("slice", SliceValidatorRequiredTestFunc([][]int{{1}}, [][]int{}))

	t.Run("map", SliceValidatorRequiredTestFunc([]map[int]int{{1: 1}}, []map[int]int{}))

	type X struct {
		A int
		B string
	}
	t.Run("struct", SliceValidatorRequiredTestFunc([]X{{A: 0, B: ""}}, []X{}))
}

func SliceValidatorRequiredTestFunc[T any, Slice []T](ok, fail Slice) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Slice[T]().Required().Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Slice[T]().Required().Build(fail).Validate(ctx)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !goval.IsCodeEqual(exp.Code, goval.SliceRequired) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.SliceRequired, exp.Code)
		}

		inp, ok := exp.Input.(Slice)
		if !ok {
			t.Fatalf("expect the error input type: %T; got error input: %T", fail, exp.Input)
		}

		if !reflect.DeepEqual(inp, fail) {
			t.Errorf("expect the error input value: %v; got error input value: %v", fail, inp)
		}

		if exp.Args != nil {
			t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
		}
	}
}

func TestSliceValidator_Min(t *testing.T) {
	t.Run("int", SliceValidatorMinTestFunc([]int{1}, []int{}))
	t.Run("int8", SliceValidatorMinTestFunc([]int8{1}, nil))
	t.Run("int16", SliceValidatorMinTestFunc([]int16{1}, []int16{}))
	t.Run("int32", SliceValidatorMinTestFunc([]int32{1}, nil))
	t.Run("int64", SliceValidatorMinTestFunc([]int64{1}, nil))

	t.Run("uint", SliceValidatorMinTestFunc([]uint{1}, nil))
	t.Run("uint8", SliceValidatorMinTestFunc([]uint8{1}, []uint8{}))
	t.Run("uint16", SliceValidatorMinTestFunc([]uint16{1}, nil))
	t.Run("uint32", SliceValidatorMinTestFunc([]uint32{1}, []uint32{}))
	t.Run("uint64", SliceValidatorMinTestFunc([]uint64{1}, nil))

	t.Run("float32", SliceValidatorMinTestFunc([]float32{1.123}, nil))
	t.Run("float64", SliceValidatorMinTestFunc([]float64{1.123}, []float64{}))

	t.Run("string", SliceValidatorMinTestFunc([]string{"abc"}, []string{}))

	t.Run("byte", SliceValidatorMinTestFunc([]byte{0x1}, []byte{}))

	t.Run("slice", SliceValidatorMinTestFunc([][]int{{1}}, [][]int{}))

	t.Run("map", SliceValidatorMinTestFunc([]map[int]int{{1: 1}}, []map[int]int{}))

	type X struct {
		A int
		B string
	}
	t.Run("struct", SliceValidatorMinTestFunc([]X{{A: 0, B: ""}}, []X{}))
}

func SliceValidatorMinTestFunc[T any, Slice []T](ok, fail Slice) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Slice[T]().Min(1).Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Slice[T]().Min(1).Build(fail).Validate(ctx)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !goval.IsCodeEqual(exp.Code, goval.SliceMin) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.SliceMin, exp.Code)
		}

		inp, ok := exp.Input.(Slice)
		if !ok {
			t.Fatalf("expect the error input type: %T; got error input: %T", fail, exp.Input)
		}

		if !reflect.DeepEqual(inp, fail) {
			t.Errorf("expect the error input value: %v; got error input value: %v", fail, inp)
		}

		args := []any{1}
		if !reflect.DeepEqual(exp.Args, args) {
			t.Errorf("expect the error args: %v, type: %T; got error args: %v, type: %T", args, args, exp.Args, exp.Args)
		}
	}
}

func TestSliceValidator_Max(t *testing.T) {
	t.Run("int", SliceValidatorMaxTestFunc([]int{1}, []int{1, 2}))
	t.Run("int8", SliceValidatorMaxTestFunc([]int8{1}, []int8{1, 2, 3}))
	t.Run("int16", SliceValidatorMaxTestFunc([]int16{1}, []int16{1, 2}))
	t.Run("int32", SliceValidatorMaxTestFunc([]int32{1}, []int32{1, 2, 3}))
	t.Run("int64", SliceValidatorMaxTestFunc([]int64{1}, []int64{1, 2, 3}))

	t.Run("uint", SliceValidatorMaxTestFunc([]uint{1}, []uint{1, 2}))
	t.Run("uint8", SliceValidatorMaxTestFunc([]uint8{1}, []uint8{1, 2}))
	t.Run("uint16", SliceValidatorMaxTestFunc([]uint16{1}, []uint16{1, 2}))
	t.Run("uint32", SliceValidatorMaxTestFunc([]uint32{1}, []uint32{1, 2}))
	t.Run("uint64", SliceValidatorMaxTestFunc([]uint64{1}, []uint64{1, 2}))

	t.Run("float32", SliceValidatorMaxTestFunc([]float32{1.123}, []float32{1, 2}))
	t.Run("float64", SliceValidatorMaxTestFunc([]float64{1.123}, []float64{1, 2}))

	t.Run("string", SliceValidatorMaxTestFunc([]string{"abc"}, []string{"a", "b"}))

	t.Run("byte", SliceValidatorMaxTestFunc([]byte{0x1}, []byte{0x1, 0x2}))

	t.Run("slice", SliceValidatorMaxTestFunc([][]int{{1}}, [][]int{{1}, {2}}))

	t.Run("map", SliceValidatorMaxTestFunc([]map[int]int{{1: 1}}, []map[int]int{{1: 1}, {2: 2}}))

	type X struct {
		A int
		B string
	}
	t.Run("struct", SliceValidatorMaxTestFunc([]X{{A: 0, B: ""}}, []X{{}, {}}))
}

func SliceValidatorMaxTestFunc[T any, Slice []T](ok, fail Slice) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		err := goval.Slice[T]().Max(1).Build(ok).Validate(ctx)
		if err != nil {
			t.Errorf("expect no error; got error: %v", err)
		}

		err = goval.Slice[T]().Max(1).Build(fail).Validate(ctx)
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !goval.IsCodeEqual(exp.Code, goval.SliceMax) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.SliceMax, exp.Code)
		}

		inp, ok := exp.Input.(Slice)
		if !ok {
			t.Fatalf("expect the error input type: %T; got error input: %T", fail, exp.Input)
		}

		if !reflect.DeepEqual(inp, fail) {
			t.Errorf("expect the error input value: %v; got error input value: %v", fail, inp)
		}

		args := []any{1}
		if !reflect.DeepEqual(exp.Args, args) {
			t.Errorf("expect the error args: %v, type: %T; got error args: %v, type: %T", args, args, exp.Args, exp.Args)
		}
	}
}

func TestSliceValidator_Each(t *testing.T) {
	ctx := context.Background()
	val := []string{"ab", "cd", "ef"}
	err := goval.Slice[string]().Each(goval.String().Min(2)).Build(val).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	val = []string{"a", "c", "e"}
	err = goval.Slice[string]().Each(goval.String().Min(2)).Build(val).Validate(ctx)
	var errs *goval.RuleErrors
	if !errors.As(err, &errs) {
		t.Fatalf("expect error type: %T; got error type: %T", errs, err)
	}

	if len(errs.Errs) != 3 {
		t.Errorf("expect number of errors: %v; got number of errors: %v", len(val), len(errs.Errs))
	}

	for i := range errs.Errs {
		if !goval.IsCodeEqual(errs.Errs[i].Code, goval.StringMin) {
			t.Errorf("errs[%d]: expect the error code: %v; got error code: %v", i, goval.StringMin, errs.Errs[i].Code)
		}

		inp, ok := errs.Errs[i].Input.(string)
		if !ok {
			t.Fatalf("errs[%d]: expect the error input type: %T; got error input: %T", i, val[i], errs.Errs[i].Input)
		}

		if inp != val[i] {
			t.Errorf("errs[%d]: expect the error input value: %v; got error input value: %v", i, val[i], inp)
		}
	}
}

func BenchmarkSliceValidator_Required(b *testing.B) {
	b.Run("int", SliceValidatorRequiredBenchmarkFunc([]int{1, 2, 3, 4}))
	b.Run("int8", SliceValidatorRequiredBenchmarkFunc([]int8{1, 2, 3, 4}))
	b.Run("int16", SliceValidatorRequiredBenchmarkFunc([]int16{1, 2, 3, 4}))
	b.Run("int32", SliceValidatorRequiredBenchmarkFunc([]int32{1, 2, 3, 4}))
	b.Run("int64", SliceValidatorRequiredBenchmarkFunc([]int64{1, 2, 3, 4}))

	b.Run("uint", SliceValidatorRequiredBenchmarkFunc([]uint{1, 2, 3, 4}))
	b.Run("uint8", SliceValidatorRequiredBenchmarkFunc([]uint8{1, 2, 3, 4}))
	b.Run("uint16", SliceValidatorRequiredBenchmarkFunc([]uint16{1, 2, 3, 4}))
	b.Run("uint32", SliceValidatorRequiredBenchmarkFunc([]uint32{1, 2, 3, 4}))
	b.Run("uint64", SliceValidatorRequiredBenchmarkFunc([]uint64{1, 2, 3, 4}))

	b.Run("float32", SliceValidatorRequiredBenchmarkFunc([]float32{1.1, 2.2, 3.3, 4.4}))
	b.Run("float64", SliceValidatorRequiredBenchmarkFunc([]float64{1.1, 2.2, 3.3, 4.4}))

	b.Run("byte", SliceValidatorRequiredBenchmarkFunc([]byte{0x1, 0x2, 0x3, 0x4}))

	b.Run("string", SliceValidatorRequiredBenchmarkFunc([]string{"1", "2", "3", "4"}))

	b.Run("slice", SliceValidatorRequiredBenchmarkFunc([][]int{{1}, {2}, {3}, {4}}))

	b.Run("map", SliceValidatorRequiredBenchmarkFunc([]map[int]int{{1: 1}, {2: 2}, {3: 3}, {4: 4}}))
}

func SliceValidatorRequiredBenchmarkFunc[T any, Slice []T](slice Slice) func(b *testing.B) {
	return func(b *testing.B) {
		ctx := context.Background()

		b.Run("empty", func(b *testing.B) {
			v := goval.Slice[T, Slice]().Required()

			for i := 0; i < b.N; i++ {
				_ = v.Build(Slice{}).Validate(ctx)
			}
		})

		b.Run("not empty", func(b *testing.B) {
			v := goval.Slice[T, Slice]().Required()

			for i := 0; i < b.N; i++ {
				_ = v.Build(slice).Validate(ctx)
			}
		})

	}
}

func BenchmarkSliceValidator_Min(b *testing.B) {
	b.Run("int", SliceValidatorMinBenchmarkFunc([]int{1, 2, 3, 4}))
	b.Run("int8", SliceValidatorMinBenchmarkFunc([]int8{1, 2, 3, 4}))
	b.Run("int16", SliceValidatorMinBenchmarkFunc([]int16{1, 2, 3, 4}))
	b.Run("int32", SliceValidatorMinBenchmarkFunc([]int32{1, 2, 3, 4}))
	b.Run("int64", SliceValidatorMinBenchmarkFunc([]int64{1, 2, 3, 4}))

	b.Run("uint", SliceValidatorMinBenchmarkFunc([]uint{1, 2, 3, 4}))
	b.Run("uint8", SliceValidatorMinBenchmarkFunc([]uint8{1, 2, 3, 4}))
	b.Run("uint16", SliceValidatorMinBenchmarkFunc([]uint16{1, 2, 3, 4}))
	b.Run("uint32", SliceValidatorMinBenchmarkFunc([]uint32{1, 2, 3, 4}))
	b.Run("uint64", SliceValidatorMinBenchmarkFunc([]uint64{1, 2, 3, 4}))

	b.Run("float32", SliceValidatorMinBenchmarkFunc([]float32{1.1, 2.2, 3.3, 4.4}))
	b.Run("float64", SliceValidatorMinBenchmarkFunc([]float64{1.1, 2.2, 3.3, 4.4}))

	b.Run("byte", SliceValidatorMinBenchmarkFunc([]byte{0x1, 0x2, 0x3, 0x4}))

	b.Run("string", SliceValidatorMinBenchmarkFunc([]string{"1", "2", "3", "4"}))

	b.Run("slice", SliceValidatorMinBenchmarkFunc([][]int{{1}, {2}, {3}, {4}}))

	b.Run("map", SliceValidatorMinBenchmarkFunc([]map[int]int{{1: 1}, {2: 2}, {3: 3}, {4: 4}}))
}

func SliceValidatorMinBenchmarkFunc[T any, Slice []T](slice Slice) func(b *testing.B) {
	return func(b *testing.B) {
		ctx := context.Background()
		n := len(slice)

		b.Run("below minimum", func(b *testing.B) {
			v := goval.Slice[T, Slice]().Min(n + 1)

			for i := 0; i < b.N; i++ {
				_ = v.Build(slice).Validate(ctx)
			}
		})

		b.Run("above minimum", func(b *testing.B) {
			v := goval.Slice[T, Slice]().Min(n - 1)

			for i := 0; i < b.N; i++ {
				_ = v.Build(slice).Validate(ctx)
			}
		})

	}
}

func BenchmarkSliceValidator_Max(b *testing.B) {
	b.Run("int", SliceValidatorMaxBenchmarkFunc([]int{1, 2, 3, 4}))
	b.Run("int8", SliceValidatorMaxBenchmarkFunc([]int8{1, 2, 3, 4}))
	b.Run("int16", SliceValidatorMaxBenchmarkFunc([]int16{1, 2, 3, 4}))
	b.Run("int32", SliceValidatorMaxBenchmarkFunc([]int32{1, 2, 3, 4}))
	b.Run("int64", SliceValidatorMaxBenchmarkFunc([]int64{1, 2, 3, 4}))

	b.Run("uint", SliceValidatorMaxBenchmarkFunc([]uint{1, 2, 3, 4}))
	b.Run("uint8", SliceValidatorMaxBenchmarkFunc([]uint8{1, 2, 3, 4}))
	b.Run("uint16", SliceValidatorMaxBenchmarkFunc([]uint16{1, 2, 3, 4}))
	b.Run("uint32", SliceValidatorMaxBenchmarkFunc([]uint32{1, 2, 3, 4}))
	b.Run("uint64", SliceValidatorMaxBenchmarkFunc([]uint64{1, 2, 3, 4}))

	b.Run("float32", SliceValidatorMaxBenchmarkFunc([]float32{1.1, 2.2, 3.3, 4.4}))
	b.Run("float64", SliceValidatorMaxBenchmarkFunc([]float64{1.1, 2.2, 3.3, 4.4}))

	b.Run("byte", SliceValidatorMaxBenchmarkFunc([]byte{0x1, 0x2, 0x3, 0x4}))

	b.Run("string", SliceValidatorMaxBenchmarkFunc([]string{"1", "2", "3", "4"}))

	b.Run("slice", SliceValidatorMaxBenchmarkFunc([][]int{{1}, {2}, {3}, {4}}))

	b.Run("map", SliceValidatorMaxBenchmarkFunc([]map[int]int{{1: 1}, {2: 2}, {3: 3}, {4: 4}}))
}

func SliceValidatorMaxBenchmarkFunc[T any, Slice []T](slice Slice) func(b *testing.B) {
	return func(b *testing.B) {
		ctx := context.Background()
		n := len(slice)

		b.Run("below maximum", func(b *testing.B) {
			v := goval.Slice[T, Slice]().Max(n + 1)

			for i := 0; i < b.N; i++ {
				_ = v.Build(slice).Validate(ctx)
			}
		})

		b.Run("above maximum", func(b *testing.B) {
			v := goval.Slice[T, Slice]().Max(n - 1)

			for i := 0; i < b.N; i++ {
				_ = v.Build(slice).Validate(ctx)
			}
		})

	}
}

func BenchmarkSliceValidator_Each(b *testing.B) {
	b.Run("int", SliceValidatorEachBenchmarkFunc[int, []int]([]int{1, 2, 3}, []int{-2, 0, 4, 5}, goval.Number[int]().Min(0).Max(3)))
	b.Run("int8", SliceValidatorEachBenchmarkFunc[int8, []int8]([]int8{1, 2, 3}, []int8{-2, 0, 4, 5}, goval.Number[int8]().Min(0).Max(3)))
	b.Run("int16", SliceValidatorEachBenchmarkFunc[int16, []int16]([]int16{1, 2, 3}, []int16{-2, 0, 4, 5}, goval.Number[int16]().Min(0).Max(3)))
	b.Run("int32", SliceValidatorEachBenchmarkFunc[int32, []int32]([]int32{1, 2, 3}, []int32{-2, 0, 4, 5}, goval.Number[int32]().Min(0).Max(3)))
	b.Run("int64", SliceValidatorEachBenchmarkFunc[int64, []int64]([]int64{1, 2, 3}, []int64{-2, 0, 4, 5}, goval.Number[int64]().Min(0).Max(3)))

	b.Run("uint", SliceValidatorEachBenchmarkFunc[uint, []uint]([]uint{1, 2, 3}, []uint{0, 1, 4, 5}, goval.Number[uint]().Min(1).Max(3)))
	b.Run("uint8", SliceValidatorEachBenchmarkFunc[uint8, []uint8]([]uint8{1, 2, 3}, []uint8{0, 1, 4, 5}, goval.Number[uint8]().Min(1).Max(3)))
	b.Run("uint16", SliceValidatorEachBenchmarkFunc[uint16, []uint16]([]uint16{1, 2, 3}, []uint16{0, 1, 4, 5}, goval.Number[uint16]().Min(1).Max(3)))
	b.Run("uint32", SliceValidatorEachBenchmarkFunc[uint32, []uint32]([]uint32{1, 2, 3}, []uint32{0, 1, 4, 5}, goval.Number[uint32]().Min(1).Max(3)))
	b.Run("uint64", SliceValidatorEachBenchmarkFunc[uint64, []uint64]([]uint64{1, 2, 3}, []uint64{0, 1, 4, 5}, goval.Number[uint64]().Min(1).Max(3)))

	b.Run("float32", SliceValidatorEachBenchmarkFunc[float32, []float32]([]float32{1.1, 2.2, 3.3}, []float32{0.5, 1.1, 4.4, 5.5}, goval.Number[float32]().Min(1).Max(3)))
	b.Run("float64", SliceValidatorEachBenchmarkFunc[float64, []float64]([]float64{1.1, 2.2, 3.3}, []float64{0.5, 1.1, 4.4, 5.5}, goval.Number[float64]().Min(1).Max(3)))

	b.Run("byte", SliceValidatorEachBenchmarkFunc[byte, []byte]([]byte{0x0, 0x1, 0x2, 0x3}, []byte{0x0, 0x1, 0x4, 0x5}, goval.Number[byte]().Min(0x1).Max(0x3)))

	b.Run("string", SliceValidatorEachBenchmarkFunc[string, []string]([]string{"a", "b", "c", "d"}, []string{"!", "1", "#", "4"}, goval.String().Match(govalregex.AlphaNumeric)))

	b.Run("slice", SliceValidatorEachBenchmarkFunc[[]int, [][]int]([][]int{{1, 2}, {1, 2}}, [][]int{{1}, {2}}, goval.Slice[int]().Min(2)))

	b.Run("map", SliceValidatorEachBenchmarkFunc[map[int]int, []map[int]int]([]map[int]int{{1: 1, 2: 2}, {3: 3, 4: 4}}, []map[int]int{{1: 1}, {2: 2}}, goval.Map[int, int]().Min(2)))
}

func SliceValidatorEachBenchmarkFunc[T any, Slice []T](validSlice Slice, errSlice Slice, validator goval.Builder[T]) func(b *testing.B) {
	return func(b *testing.B) {
		ctx := context.Background()

		v := goval.Slice[T, Slice]().Each(validator)

		b.Run("valid slice", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = v.Build(validSlice).Validate(ctx)
			}
		})

		b.Run("error slice", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = v.Build(errSlice).Validate(ctx)
			}
		})
	}
}
