package goval_test

import (
	"context"
	"errors"
	"github.com/pkg-id/goval"
	"testing"
)

func TestStringValidator_Required(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		err := goval.String().Required().WithValue("").Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		err = goval.String().Required().WithValue("   ").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.String().Required().WithValue("a").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("named", func(t *testing.T) {
		err := goval.Named("name", goval.String().Required().WithValue("")).Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.KeyError
		if !errors.As(err, &expErr) {
			t.Fatalf("expect KeyError")
		}

		if expErr.Key != "name" {
			t.Errorf("expect key of KeyError is name, but got %v", expErr.Key)
		}

		err = goval.Named("name", goval.String().Required().WithValue("   ")).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Named("name", goval.String().Required().WithValue("a")).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.String().Required().WithValue(""))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Errorf("expect Errors")
		}

		err = goval.Execute(context.Background(), goval.String().Required().WithValue("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Required().WithValue("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped-and-named", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.Named("name", goval.String().Required().WithValue("")))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Fatalf("expect Errors")
		}

		var keyErr *goval.KeyError
		if !errors.As(expErr.Errs()[0], &keyErr) {
			t.Fatalf("expect KeyError")
		}

		if keyErr.Key != "name" {
			t.Errorf("expect key of KeyError is name, but got %v", keyErr.Key)
		}

		err = goval.Execute(context.Background(), goval.String().Required().WithValue("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Required().WithValue("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func TestStringValidator_Min(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		err := goval.String().Min(1).WithValue("").Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		err = goval.String().Min(1).WithValue("   ").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.String().Min(1).WithValue("a").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("named", func(t *testing.T) {
		err := goval.Named("name", goval.String().Min(1).WithValue("")).Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.KeyError
		if !errors.As(err, &expErr) {
			t.Fatalf("expect KeyError")
		}

		if expErr.Key != "name" {
			t.Errorf("expect key of KeyError is name, but got %v", expErr.Key)
		}

		err = goval.Named("name", goval.String().Min(1).WithValue("   ")).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Named("name", goval.String().Min(1).WithValue("a")).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.String().Min(1).WithValue(""))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Errorf("expect Errors")
		}

		err = goval.Execute(context.Background(), goval.String().Min(1).WithValue("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Min(1).WithValue("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped-and-named", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.Named("name", goval.String().Min(1).WithValue("")))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Fatalf("expect Errors")
		}

		var keyErr *goval.KeyError
		if !errors.As(expErr.Errs()[0], &keyErr) {
			t.Fatalf("expect KeyError")
		}

		if keyErr.Key != "name" {
			t.Errorf("expect key of KeyError is name, but got %v", keyErr.Key)
		}

		err = goval.Execute(context.Background(), goval.String().Min(1).WithValue("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Min(1).WithValue("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func TestStringValidator_Max(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		err := goval.String().Max(1).WithValue("ab").Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		err = goval.String().Max(5).WithValue("   ").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.String().Max(1).WithValue("a").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("named", func(t *testing.T) {
		err := goval.Named("name", goval.String().Max(1).WithValue("ab")).Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.KeyError
		if !errors.As(err, &expErr) {
			t.Fatalf("expect KeyError")
		}

		if expErr.Key != "name" {
			t.Errorf("expect key of KeyError is name, but got %v", expErr.Key)
		}

		err = goval.Named("name", goval.String().Max(5).WithValue("   ")).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Named("name", goval.String().Max(1).WithValue("a")).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.String().Max(1).WithValue("ab"))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Errorf("expect Errors")
		}

		err = goval.Execute(context.Background(), goval.String().Max(5).WithValue("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Max(1).WithValue("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped-and-named", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.Named("name", goval.String().Max(1).WithValue("ab")))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Fatalf("expect Errors")
		}

		var keyErr *goval.KeyError
		if !errors.As(expErr.Errs()[0], &keyErr) {
			t.Fatalf("expect KeyError")
		}

		if keyErr.Key != "name" {
			t.Errorf("expect key of KeyError is name, but got %v", keyErr.Key)
		}

		err = goval.Execute(context.Background(), goval.String().Max(5).WithValue("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Max(1).WithValue("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func TestString(t *testing.T) {
	rules := goval.String().Required().Min(2).Max(5)

	err := rules.WithValue("").Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error since the required rules is violeted")
	}

	err = rules.WithValue("a").Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error since the min length rules is violeted")
	}

	err = rules.WithValue("abc abc").Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error since the max length rules is violeted")
	}

	err = rules.WithValue("abc").Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error since no rule is violated. but got error: %v", err)
	}
}
