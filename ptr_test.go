package goval_test

import (
	"context"
	"errors"
	"testing"

	"github.com/pkg-id/goval"
)

func TestPtr(t *testing.T) {
	ctx := context.Background()
	err := goval.Ptr[string]().Build(nil).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestPtrValidator_Required(t *testing.T) {
	ctx := context.Background()
	err := goval.Ptr[string]().Required().Build(nil).Validate(ctx)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !goval.IsCodeEqual(exp.Code, goval.NilRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.NilRequired, exp.Code)
	}

	inp, ok := exp.Input.(*string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", (*string)(nil), exp.Input)
	}

	if inp != nil {
		t.Errorf("expect the error input value: %v; got error input value: %v", nil, inp)
	}

	if exp.Args != nil {
		t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
	}

	err = goval.Ptr[string]().Required().Build(new(string)).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestPtrValidator_Optional(t *testing.T) {
	ctx := context.Background()
	sv := goval.String().Required()
	err := goval.Ptr[string]().Optional(sv).Build(nil).Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.Ptr[string]().Optional(sv).Build(new(string)).Validate(ctx)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !goval.IsCodeEqual(exp.Code, goval.StringRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringRequired, exp.Code)
	}

	inp, ok := exp.Input.(string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", "", exp.Input)
	}

	if inp != "" {
		t.Errorf("expect the error input value: %q; got error input value: %q", "", inp)
	}

	if exp.Args != nil {
		t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
	}
}

func TestPtrValidator_Then(t *testing.T) {
	ctx := context.Background()
	sv := goval.String().Required()
	err := goval.Ptr[string]().Then(sv).Build(new(string)).Validate(ctx)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !goval.IsCodeEqual(exp.Code, goval.StringRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringRequired, exp.Code)
	}

	inp, ok := exp.Input.(string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", "", exp.Input)
	}

	if inp != "" {
		t.Errorf("expect the error input value: %q; got error input value: %q", "", inp)
	}

	if exp.Args != nil {
		t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
	}
}

func TestPtrValidator_ThenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect panic; got no panic")
		}
	}()

	sv := goval.String().Required()
	goval.Ptr[string]().Then(sv).Build(nil).Validate(context.Background())
}
