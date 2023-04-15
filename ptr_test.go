package goval_test

import (
	"context"
	"errors"
	"testing"

	"github.com/pkg-id/goval"
)

func TestPtr(t *testing.T) {
	ctx := context.Background()
	err := goval.Ptr[string]().Validate(ctx, nil)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestPtrValidator_Required(t *testing.T) {
	ctx := context.Background()
	err := goval.Ptr[string]().Required().Validate(ctx, nil)
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.PtrRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.PtrRequired, exp.Code)
	}

	if exp.Args != nil {
		t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
	}

	err = goval.Ptr[string]().Required().Validate(ctx, new(string))
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestPtrValidator_Optional(t *testing.T) {
	ctx := context.Background()
	sv := goval.String().Required()
	err := goval.Ptr[string]().Optional(sv).Validate(ctx, nil)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.Ptr[string]().Optional(sv).Validate(ctx, new(string))
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.StringRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringRequired, exp.Code)
	}

	if exp.Args != nil {
		t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
	}
}

func TestPtrValidator_Then(t *testing.T) {
	ctx := context.Background()
	sv := goval.String().Required()
	err := goval.Ptr[string]().Then(sv).Validate(ctx, new(string))
	if err == nil {
		t.Errorf("expect error; got no error")
	}

	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.StringRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringRequired, exp.Code)
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
	_ = goval.Ptr[string]().Then(sv).Validate(context.Background(), nil)
}

func BenchmarkPtrValidator_Required(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		for i := 0; i < b.N; i++ {
			_ = goval.Ptr[int]().Required().Validate(ctx, nil)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 1
		for i := 0; i < b.N; i++ {
			_ = goval.Ptr[int]().Required().Validate(ctx, &val)
		}
	})
}

func BenchmarkPtrValidator_Optional(b *testing.B) {
	b.Run("when rules violated", func(b *testing.B) {
		ctx := context.Background()
		for i := 0; i < b.N; i++ {
			_ = goval.Ptr[int]().Optional(goval.Number[int]().Required()).Validate(ctx, nil)
		}
	})

	b.Run("when no rules violated", func(b *testing.B) {
		ctx := context.Background()
		val := 1
		for i := 0; i < b.N; i++ {
			_ = goval.Ptr[int]().Optional(goval.Number[int]().Required()).Validate(ctx, &val)
		}
	})
}
