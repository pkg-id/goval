package goval_test

import (
	"context"
	"github.com/pkg-id/goval"
	"testing"
)

func TestSliceValidator_Required(t *testing.T) {
	err := goval.Slice[int]().Required().Build(nil).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Slice[int]().Required().Build([]int{}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Slice[int]().Required().Build([]int{1, 2}).Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}

	err = goval.Slice[string]().Required().Build(nil).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Slice[string]().Required().Build([]string{}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Slice[string]().Required().Build([]string{"a", "b"}).Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}
}

func TestSliceValidator_Min(t *testing.T) {
	err := goval.Slice[int]().Required().Min(3).Build([]int{1, 2}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Slice[int]().Required().Min(1).Build([]int{1, 2}).Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}
}

func TestSliceValidator_Max(t *testing.T) {
	err := goval.Slice[int]().Required().Min(1).Max(1).Build([]int{1, 2}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Slice[int]().Required().Min(1).Max(10).Build([]int{1, 2}).Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}
}

func TestSliceValidator_Each(t *testing.T) {
	err := goval.Slice[int]().Required().Min(1).Max(3).
		Each(goval.Number[int]().Min(4)).
		Build([]int{1, 2}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Slice[int]().Required().Min(1).Max(3).
		Each(goval.Number[int]().Min(4)).
		Build([]int{10, 20}).Validate(context.Background())

	if err != nil {
		t.Errorf("expect no error")
	}
}
