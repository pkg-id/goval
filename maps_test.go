package goval_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/pkg-id/goval"
)

func TestMap(t *testing.T) {
	ctx := context.Background()
	err := goval.Map[string, string]().Build(nil).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestMapValidator_Required(t *testing.T) {
	ctx := context.Background()
	err := goval.Map[string, string]().Required().Build(map[string]string{}).Validate(ctx)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.MapRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.MapRequired, exp.Code)
	}

	inp, ok := exp.Input.(map[string]string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", map[string]string{}, exp.Input)
	}

	if len(inp) != 0 {
		t.Errorf("expect the error input value: %v; got error input value: %v", map[string]string{}, inp)
	}

	if exp.Args != nil {
		t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
	}

	err = goval.Map[string, string]().Required().Build(map[string]string{"key": "value"}).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestMapValidator_Min(t *testing.T) {
	ctx := context.Background()
	val := map[string]string{"key": "value"}
	err := goval.Map[string, string]().Min(2).Build(val).Validate(ctx)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.MapMin) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.MapMin, exp.Code)
	}

	inp, ok := exp.Input.(map[string]string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", val, exp.Input)
	}

	if len(inp) != len(val) {
		t.Errorf("expect the error input value: %v; got error input value: %v", val, inp)
	}

	args := []any{2}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}

	err = goval.Map[string, string]().Min(1).Build(val).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestMapValidator_Max(t *testing.T) {
	ctx := context.Background()
	val := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	err := goval.Map[string, string]().Max(2).Build(val).Validate(ctx)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.MapMax) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.MapMax, exp.Code)
	}

	inp, ok := exp.Input.(map[string]string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", val, exp.Input)
	}

	if len(inp) != len(val) {
		t.Errorf("expect the error input value: %v; got error input value: %v", val, inp)
	}

	args := []any{2}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}

	err = goval.Map[string, string]().Max(3).Build(val).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestMapValidator_Each(t *testing.T) {
	ctx := context.Background()
	val := map[string]string{
		"key1": "",
		"key2": "a",
	}

	sv := goval.String().Required().Min(2)
	err := goval.Map[string, string]().Each(sv).Build(val).Validate(ctx)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp goval.Errors
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if len(exp) != 2 {
		t.Errorf("expect the error length: %d; got error length: %d", 2, len(exp))
	}
}
