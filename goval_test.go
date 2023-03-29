package goval_test

import (
	"context"
	"errors"
	"github.com/pkg-id/goval"
	"testing"
)

func TestNamed(t *testing.T) {
	t.Run("when validation fails", func(t *testing.T) {
		ctx := context.Background()
		err := goval.Named("field-name", "", goval.String().Required()).Validate(ctx)

		var exp *goval.KeyError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if exp.Key != "field-name" {
			t.Errorf("expect error key is field-name; got %s", exp.Key)
		}

		if exp.Err == nil {
			t.Errorf("expect error is not nil")
		}
	})

	t.Run("when validation ok", func(t *testing.T) {
		ctx := context.Background()
		err := goval.Named("field-name", "a", goval.String().Required()).Validate(ctx)
		if err != nil {
			t.Fatalf("expect not error")
		}
	})
}

func TestEach(t *testing.T) {
	t.Run("when validation fails", func(t *testing.T) {
		ctx := context.Background()
		val := []string{"a", "bc", "d", "ef"}

		err := goval.Each(goval.String().Required().Min(2).Build).Build(val).Validate(ctx)
		var exp goval.Errors
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if len(exp) != 2 {
			t.Fatalf("expect two errors")
		}
	})

	t.Run("when validation ok", func(t *testing.T) {
		ctx := context.Background()
		val := []string{"aa", "bc", "dd", "ef"}

		err := goval.Each(goval.String().Required().Min(2).Build).Build(val).Validate(ctx)
		if err != nil {
			t.Fatalf("expect not error")
		}
	})
}

func TestExecute(t *testing.T) {
	t.Run("when validation fails", func(t *testing.T) {
		ctx := context.Background()

		err := goval.Execute(ctx,
			goval.String().Required().Min(2).Build("a"),
			goval.Number[int]().Required().Min(8).Build(7),
		)
		var exp goval.Errors
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if len(exp) != 2 {
			t.Fatalf("expect two errors")
		}
	})

	t.Run("when validation ok", func(t *testing.T) {
		ctx := context.Background()

		err := goval.Execute(ctx,
			goval.String().Required().Min(2).Build("ab"),
			goval.Number[int]().Required().Min(8).Build(8),
		)
		if err != nil {
			t.Fatalf("expect not error")
		}
	})
}
