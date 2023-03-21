package goval_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/pkg-id/goval"
	"github.com/pkg-id/goval/govalregex"
)

func TestString(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Build("").Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestStringValidator_Required(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Required().Build("abc").Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Required().Build("").Validate(ctx)
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

func TestStringValidator_Min(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Min(3).Build("abc").Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Min(3).Build("ab").Validate(ctx)
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !goval.IsCodeEqual(exp.Code, goval.StringMin) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringMin, exp.Code)
	}

	inp, ok := exp.Input.(string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", "", exp.Input)
	}

	if inp != "ab" {
		t.Errorf("expect the error input value: %q; got error input value: %q", "", inp)
	}

	args := []any{3}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestStringValidator_Max(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Max(3).Build("abc").Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Max(2).Build("abc").Validate(ctx)
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !goval.IsCodeEqual(exp.Code, goval.StringMax) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringMax, exp.Code)
	}

	inp, ok := exp.Input.(string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", "", exp.Input)
	}

	if inp != "abc" {
		t.Errorf("expect the error input value: %q; got error input value: %q", "", inp)
	}

	args := []any{2}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestStringValidator_Match(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Match(govalregex.AlphaNumeric).Build("abc123").Validate(ctx)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Match(govalregex.AlphaNumeric).Build("abc??").Validate(ctx)
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !goval.IsCodeEqual(exp.Code, goval.StringMatch) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringMatch, exp.Code)
	}

	inp, ok := exp.Input.(string)
	if !ok {
		t.Fatalf("expect the error input type: %T; got error input: %T", "", exp.Input)
	}

	if inp != "abc??" {
		t.Errorf("expect the error input value: %q; got error input value: %q", "", inp)
	}

	args := []any{govalregex.AlphaNumeric.RegExp().String()}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func BenchmarkStringValidator_Build(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goval.String().Build("")
	}
}

func BenchmarkStringValidator_Required(b *testing.B) {
	ctx := context.Background()

	b.Run("benchmark without value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Required().Build("").Validate(ctx)
		}
	})

	b.Run("benchmark with value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Required().Build("random").Validate(ctx)
		}
	})
}

func BenchmarkStringValidator_Min(b *testing.B) {
	ctx := context.Background()

	b.Run("with value under minimum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Min(10).Build("12345").Validate(ctx)
		}
	})

	b.Run("with value above minimum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Min(5).Build("123456").Validate(ctx)
		}
	})
}

func BenchmarkStringValidator_Max(b *testing.B) {
	ctx := context.Background()
	b.Run("with value above maximum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Max(5).Build("123456").Validate(ctx)
		}
	})

	b.Run("with value under maximum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Max(5).Build("1234").Validate(ctx)
		}
	})
}

func BenchmarkStringValidator_Match(b *testing.B) {
	ctx := context.Background()

	emailRegex := govalregex.NewLazy("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")
	ipRegex := govalregex.NewLazy("(?:(?:2(?:[0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9])\\.){3}(?:(?:2([0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9]))")

	b.Run("email regex validation with valid email", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Match(emailRegex).Build("email@example.com").Validate(ctx)
		}
	})

	b.Run("email regex validation with invalid email", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Match(emailRegex).Build("emailexample.com").Validate(ctx)
		}
	})

	b.Run("ip address regex validation with valid ip", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Match(ipRegex).Build("127.0.0.1").Validate(ctx)
		}
	})

	b.Run("ip address regex validation with invalid ip", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = goval.String().Match(ipRegex).Build("a.b.c.d").Validate(ctx)
		}
	})

}
