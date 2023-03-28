package goval_test

import (
	"errors"
	"github.com/pkg-id/goval"
	"testing"
)

func TestRuleError(t *testing.T) {
	err := goval.NewRuleError(goval.NumberMin, 2, 3)
	exp := `{"code":3001,"input":2,"args":[3]}`
	got := err.Error()
	if got != exp {
		t.Errorf("expect string Error: %q; got %q", exp, got)
	}
}

func TestTextError(t *testing.T) {
	err := goval.TextError("my-error")
	exp := "my-error"
	got := err.Error()
	if got != exp {
		t.Errorf("expect string Error: %q; got %q", exp, got)
	}
}

func TestKeyError(t *testing.T) {
	t.Run("Error: using marshal able error", func(t *testing.T) {
		err := goval.KeyError{
			Key: "my-key",
			Err: goval.TextError("my-error"),
		}

		str := err.Error()
		exp := `{"key":"my-key","err":"my-error"}`
		if str != exp {
			t.Errorf("expect string Error: %q; got %q", exp, str)
		}
	})

	t.Run("Error: using an error that not implement json marshaller", func(t *testing.T) {
		err := goval.KeyError{
			Key: "my-key",
			Err: errors.New("my-error"), // it should be converted by the json.Marshal of KeyError.
		}

		str := err.Error()
		exp := `{"key":"my-key","err":"my-error"}`
		if str != exp {
			t.Errorf("expect string Error: %q; got %q", exp, str)
		}
	})
}

func TestErrors(t *testing.T) {
	t.Run("Error: using var", func(t *testing.T) {
		var errs goval.Errors
		str := errs.Error()
		if str != "null" {
			t.Errorf("expect string Error: %q; got %q", "null", str)
		}
	})

	t.Run("Error: using make", func(t *testing.T) {
		errs := make(goval.Errors, 0)
		str := errs.Error()
		if str != "[]" {
			t.Errorf("expect string Error: %q; got %q", "[]", str)
		}
	})
}
