package goval_test

import (
	"context"
	"errors"
	"testing"

	"github.com/pkg-id/goval"
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

	t.Run("internal error", func(t *testing.T) {
		ctx := context.Background()

		retErr := errors.New("internal error")
		custom := func(ctx context.Context, value string) error {
			return goval.NewInternalError(retErr)
		}

		err := goval.Named("field-name", "a", goval.String().With(custom)).Validate(ctx)
		if !errors.Is(err, retErr) {
			t.Fatalf("expect error is internal error")
		}
	})

	t.Run("without internal error", func(t *testing.T) {
		ctx := context.Background()

		retErr := errors.New("internal error")
		custom := func(ctx context.Context, value string) error {
			return retErr
		}

		err := goval.Named("field-name", "a", goval.String().With(custom)).Validate(ctx)
		var kErr *goval.KeyError
		if !errors.As(err, &kErr) {
			t.Fatalf("expect error type: %T; got error type: %T", kErr, err)
		}

		if !errors.Is(kErr.Err, retErr) {
			t.Fatalf("expect error is internal error")
		}
	})
}

func TestEach(t *testing.T) {
	t.Run("when validation fails", func(t *testing.T) {
		ctx := context.Background()
		val := []string{"a", "bc", "d", "ef"}

		err := goval.Each[string](goval.String().Required().Min(2)).Validate(ctx, val)
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

		err := goval.Each[string](goval.String().Required().Min(2)).Validate(ctx, val)
		if err != nil {
			t.Fatalf("expect not error")
		}
	})
}

func TestExecute(t *testing.T) {
	t.Run("when validation fails", func(t *testing.T) {
		ctx := context.Background()

		err := goval.Execute(ctx,
			goval.Bind[string]("a", goval.String().Required().Min(2)),
			goval.Bind[int](7, goval.Number[int]().Required().Min(8)),
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
			goval.Bind[string]("ab", goval.String().Required().Min(2)),
			goval.Bind[int](8, goval.Number[int]().Required().Min(8)),
		)
		if err != nil {
			t.Fatalf("expect not error")
		}
	})

	t.Run("when error is unknown type", func(t *testing.T) {
		ctx := context.Background()

		internalError := errors.New("internal error")
		customValidator := func(ctx context.Context, value int) error {
			return internalError
		}

		err := goval.Execute(ctx,
			goval.Bind[string]("a", goval.String().Required().Min(2)),
			goval.Bind[int](8, goval.Number[int]().Required().With(customValidator)),
		)

		if !errors.Is(err, internalError) {
			t.Fatalf("expect validation error is discarded and internal error is returned; but got %v", err)
		}
	})
}

func TestUse(t *testing.T) {

	type Product struct {
		ID    int64   `json:"id"`
		Price float64 `json:"price"`
	}

	validator := func(ctx context.Context, p Product) error {
		return goval.Execute(ctx,
			goval.Named("id", p.ID, goval.Number[int64]().Required()),
			goval.Named("price", p.Price, goval.Number[float64]().Required()),
		)
	}

	t.Run("when validation fails", func(t *testing.T) {
		var p Product

		ctx := context.Background()
		err := goval.Execute(ctx, goval.Named("product", p, goval.Use(validator)))

		var exp goval.Errors
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}
	})

	t.Run("when validation ok", func(t *testing.T) {
		p := Product{
			ID:    1,
			Price: 10_000,
		}

		ctx := context.Background()
		err := goval.Execute(ctx, goval.Named("product", p, goval.Use(validator)))

		if err != nil {
			t.Fatalf("expect not error")
		}
	})
}
