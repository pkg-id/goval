package goval_test

import (
	"context"
	"errors"
	"github.com/pkg-id/goval"
	"github.com/pkg-id/goval/govalregex"
	"testing"
)

func TestStringValidator_Required(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		err := goval.String().Required().Build("").Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		err = goval.String().Required().Build("   ").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.String().Required().Build("a").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("named", func(t *testing.T) {
		err := goval.Named[string]("name", "", goval.String().Required()).Validate(context.Background())
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

		err = goval.Named[string]("name", "   ", goval.String().Required()).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Named[string]("name", "a", goval.String().Required()).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.String().Required().Build(""))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Errorf("expect Errors")
		}

		err = goval.Execute(context.Background(), goval.String().Required().Build("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Required().Build("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped-and-named", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.Named[string]("name", "", goval.String().Required()))
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

		err = goval.Execute(context.Background(), goval.String().Required().Build("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Required().Build("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func TestStringValidator_Min(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		err := goval.String().Min(1).Build("").Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		err = goval.String().Min(1).Build("   ").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.String().Min(1).Build("a").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("named", func(t *testing.T) {
		err := goval.Named("name", "", goval.String().Min(1)).Validate(context.Background())
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

		err = goval.Named("name", "    ", goval.String().Min(1)).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Named("name", "a", goval.String().Min(1)).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.String().Min(1).Build(""))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Errorf("expect Errors")
		}

		err = goval.Execute(context.Background(), goval.String().Min(1).Build("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Min(1).Build("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped-and-named", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.Named("name", "", goval.String().Min(1)))
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

		err = goval.Execute(context.Background(), goval.String().Min(1).Build("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Min(1).Build("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func TestStringValidator_Max(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		err := goval.String().Max(1).Build("ab").Validate(context.Background())
		if err == nil {
			t.Errorf("expect got an error")
		}

		err = goval.String().Max(5).Build("   ").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.String().Max(1).Build("a").Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("named", func(t *testing.T) {
		err := goval.Named("name", "ab", goval.String().Max(1)).Validate(context.Background())
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

		err = goval.Named("name", "  ", goval.String().Max(5)).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Named("name", "a", goval.String().Max(1)).Validate(context.Background())
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.String().Max(1).Build("ab"))
		if err == nil {
			t.Errorf("expect got an error")
		}

		var expErr *goval.Errors
		if !errors.As(err, &expErr) {
			t.Errorf("expect Errors")
		}

		err = goval.Execute(context.Background(), goval.String().Max(5).Build("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Max(1).Build("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})

	t.Run("grouped-and-named", func(t *testing.T) {
		err := goval.Execute(context.Background(), goval.Named("name", "ab", goval.String().Max(1)))
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

		err = goval.Execute(context.Background(), goval.String().Max(5).Build("   "))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}

		err = goval.Execute(context.Background(), goval.String().Max(1).Build("a"))
		if err != nil {
			t.Errorf("expect no error but got %v", err)
		}
	})
}

func TestString(t *testing.T) {
	rules := goval.String().Required().Min(2).Max(5)

	err := rules.Build("").Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error since the required rules is violeted")
	}

	err = rules.Build("a").Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error since the min length rules is violeted")
	}

	err = rules.Build("abc abc").Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error since the max length rules is violeted")
	}

	err = rules.Build("abc").Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error since no rule is violated. but got error: %v", err)
	}
}

func TestStringValidator_Match(t *testing.T) {
	err := goval.String().Match(govalregex.AlphaNumeric).Build("#").Validate(context.Background())
	if err == nil {
		t.Errorf("expect got an error: %v", err)
	}

	err = goval.String().Match(govalregex.AlphaNumeric).Build("abc123").Validate(context.Background())
	if err != nil {
		t.Errorf("expect no error")
	}
}
