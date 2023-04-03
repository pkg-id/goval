package goval_test

import (
	"context"
	"errors"
	"github.com/pkg-id/goval"
	"reflect"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ctx := context.Background()
	err := goval.Time().Validate(ctx, time.Now())
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}
}

func TestTimeValidator_Required(t *testing.T) {
	ctx := context.Background()
	err := goval.Time().Required().Validate(ctx, time.Now())
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.Time().Required().Validate(ctx, time.Time{})
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.TimeRequired) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.TimeRequired, exp.Code)
	}

	if exp.Args != nil {
		t.Errorf("expect the error args is empty; got error args: %v", exp.Args)
	}
}

func TestTimeValidator_Min(t *testing.T) {
	ctx := context.Background()
	tNow := time.Date(2022, 01, 02, 0, 0, 0, 0, time.UTC)
	tMin := time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)

	err := goval.Time().Min(tMin).Validate(ctx, tNow)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.Time().Min(tNow).Validate(ctx, tMin)
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.TimeMin) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.TimeMin, exp.Code)
	}

	args := []any{tNow}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestTimeValidator_Max(t *testing.T) {
	ctx := context.Background()
	tNow := time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)
	tMax := time.Date(2022, 01, 02, 0, 0, 0, 0, time.UTC)

	err := goval.Time().Max(tMax).Validate(ctx, tMax)
	if err != nil {
		t.Errorf("expect no error; got error: %v", err)
	}

	err = goval.Time().Max(tNow).Validate(ctx, tMax)
	var exp *goval.RuleError
	if !errors.As(err, &exp) {
		t.Fatalf("expect error type: %T; got error type: %T", exp, err)
	}

	if !exp.Code.Equal(goval.TimeMax) {
		t.Errorf("expect the error code: %v; got error code: %v", goval.TimeMax, exp.Code)
	}

	args := []any{tNow}
	if !reflect.DeepEqual(exp.Args, args) {
		t.Errorf("expect the error args: %v; got error args: %v", args, exp.Args)
	}
}

func TestTimeValidator_When(t *testing.T) {
	jkt, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatalf("expect load ok, but got an error: %v", err)
	}

	isUTC := func(v time.Time) bool { return v.Location() == time.UTC }
	isJakarta := func(v time.Time) bool { return v.Location() == jkt }

	now := time.Now().In(time.UTC)

	validator := goval.Time().
		Required().
		When(isUTC, func(chain goval.TimeValidator) goval.TimeValidator { return chain.Min(now) }).
		When(isJakarta, func(chain goval.TimeValidator) goval.TimeValidator {
			return chain.Min(now.Add(7 * time.Hour).In(jkt))
		}).
		Max(now.Add(24 * time.Hour))

	t.Run("required", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, time.Time{})
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.TimeRequired) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.TimeRequired, exp.Code)
		}
	})

	t.Run("min when UTC", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, now.Add(-1*time.Hour))
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.TimeMin) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.TimeMin, exp.Code)
		}
	})

	t.Run("min when Asia/Jakarta", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, now.In(jkt))
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.TimeMin) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.TimeMin, exp.Code)
		}
	})

	t.Run("fails on max", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, now.In(jkt).Add(56*time.Hour))
		var exp *goval.RuleError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if !exp.Code.Equal(goval.TimeMax) {
			t.Errorf("expect the error code: %v; got error code: %v", goval.TimeMax, exp.Code)
		}
	})

	t.Run("valid", func(t *testing.T) {
		ctx := context.Background()
		err := validator.Validate(ctx, now.In(jkt).Add(8*time.Hour))
		if err != nil {
			t.Fatalf("expect no error, but got an error: %v", err)
		}
	})
}
