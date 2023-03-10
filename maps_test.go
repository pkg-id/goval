package goval_test

import (
	"context"
	"github.com/pkg-id/goval"
	"testing"
)

func TestMapValidator_Required(t *testing.T) {
	err := goval.Map[string, int]().Required().Build(nil).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Map[string, int]().Required().Build(map[string]int{}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Map[string, int]().Required().Build(map[string]int{"a": 1, "b": 2}).Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}
}

func TestMapValidator_Min(t *testing.T) {
	err := goval.Map[string, int]().Required().Min(3).Build(map[string]int{"a": 1, "b": 2}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Map[string, int]().Required().Min(1).Build(map[string]int{"a": 1, "b": 2}).Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}
}

func TestMapValidator_Max(t *testing.T) {
	err := goval.Map[string, int]().Required().Min(1).Max(1).Build(map[string]int{"a": 1, "b": 2}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Map[string, int]().Required().Min(1).Max(10).Build(map[string]int{"a": 1, "b": 2}).Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}
}

func TestMapValidator_Each(t *testing.T) {
	err := goval.Map[string, int]().Required().Min(1).Max(3).
		Each(goval.Number[int]().Min(4)).
		Build(map[string]int{"a": 1, "b": 2}).Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.Map[string, int]().Required().Min(1).Max(3).
		Each(goval.Number[int]().Min(4)).
		Build(map[string]int{"a": 10, "b": 20}).Validate(context.Background())

	if err != nil {
		t.Errorf("expect no error")
	}
}
