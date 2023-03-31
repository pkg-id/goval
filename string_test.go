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
	err := goval.String().Validate(ctx, "")
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestStringValidator_Required(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Required().Validate(ctx, "abc")
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Required().Validate(ctx, "")
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

func TestStringValidator_Min(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Min(3).Validate(ctx, "abc")
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Min(3).Validate(ctx, "ab")
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.StringMin) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringMin, exp.Code)
	}

	args := []any{3}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestStringValidator_Max(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Max(3).Validate(ctx, "abc")
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Max(2).Validate(ctx, "abc")
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.StringMax) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringMax, exp.Code)
	}

	args := []any{2}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestStringValidator_Match(t *testing.T) {
	ctx := context.Background()
	err := goval.String().Match(govalregex.AlphaNumeric).Validate(ctx, "abc123")
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().Match(govalregex.AlphaNumeric).Validate(ctx, "abc??")
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.StringMatch) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringMatch, exp.Code)
	}

	args := []any{govalregex.AlphaNumeric.RegExp().String()}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestStringValidator_In(t *testing.T) {
	ctx := context.Background()
	err := goval.String().In("a", "b", "c").Validate(ctx, "a")
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().In("a", "b", "c").Validate(ctx, "A")
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.StringIn) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringIn, exp.Code)
	}

	args := []any{[]string{"a", "b", "c"}}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestStringValidator_InFold(t *testing.T) {
	ctx := context.Background()
	err := goval.String().InFold("a", "b", "c").Validate(ctx, "C")
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.String().InFold("a", "b", "c").Validate(ctx, "Z")
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.StringInFold) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.StringInFold, exp.Code)
	}

	args := []any{[]string{"a", "b", "c"}}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func BenchmarkStringValidator_Required(b *testing.B) {
	ctx := context.Background()

	v := goval.String().Required()
	b.Run("benchmark without value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v.Validate(ctx, "")
		}
	})

	b.Run("benchmark with value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v.Validate(ctx, "random")
		}
	})
}

func BenchmarkStringValidator_Min(b *testing.B) {
	ctx := context.Background()
	v := goval.String().Min(5)

	b.Run("with value under minimum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v.Validate(ctx, "1234")
		}
	})

	b.Run("with value above minimum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v.Validate(ctx, "123456")
		}
	})
}

func BenchmarkStringValidator_Max(b *testing.B) {
	ctx := context.Background()
	v := goval.String().Max(5)

	b.Run("with value above maximum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v.Validate(ctx, "123456")
		}
	})

	b.Run("with value under maximum character", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = v.Validate(ctx, "1234")
		}
	})
}

func BenchmarkStringValidator_Match(b *testing.B) {
	ctx := context.Background()

	emailRegex := govalregex.NewLazy("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")
	ipRegex := govalregex.NewLazy("(?:(?:2(?:[0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9])\\.){3}(?:(?:2([0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9]))")

	emailValidator := goval.String().Match(emailRegex)
	ipValidator := goval.String().Match(ipRegex)

	b.Run("email regex validation with valid email", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = emailValidator.Validate(ctx, "email@example.com")
		}
	})

	b.Run("email regex validation with invalid email", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = emailValidator.Validate(ctx, "emailexample.com")
		}
	})

	b.Run("ip address regex validation with valid ip", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ipValidator.Validate(ctx, "127.0.0.1")
		}
	})

	b.Run("ip address regex validation with invalid ip", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ipValidator.Validate(ctx, "a.b.c.d")
		}
	})

}
