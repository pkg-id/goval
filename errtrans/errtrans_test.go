package errtrans

import (
	"context"
	"github.com/pkg-id/goval"
	"strings"
	"testing"
)

type ruleCode bool

func (r ruleCode) Equal(_ goval.RuleCoder) bool {
	return bool(r)
}

func TestTranslator_Translate(t *testing.T) {
	ctx := context.Background()
	bundle, _ := DefaultBundle()
	tr := NewTranslator(WithBundle(bundle))

	t.Run("when validation ok", func(t *testing.T) {
		ruleErr := goval.NewRuleError(goval.NumberRequired)

		err := tr.Translate(ctx, ruleErr)
		if err == nil {
			t.Fatalf("expect error; got nil")
		}

		if err.Error() != "This field is required." {
			t.Errorf("expect error field required; got %v", err)
		}
	})

	t.Run("when use invalid rule code", func(t *testing.T) {
		ruleErr := goval.NewRuleError(ruleCode(false))

		err := tr.Translate(ctx, ruleErr)
		if !strings.HasPrefix(err.Error(), "RuleError") {
			t.Errorf("expect RuleError; got %v", err)
		}
	})

	t.Run("when use invalid language key", func(t *testing.T) {
		ruleErr := goval.NewRuleError(goval.NumberRequired)

		ctx := ContextWithLanguage(ctx, "es")

		err := tr.Translate(ctx, ruleErr)
		if err != ErrLanguageDictionaryIsNotFound {
			t.Errorf("expect error %v; got %v", ErrLanguageDictionaryIsNotFound, err)
		}
	})

	tr = NewTranslator()
	t.Run("when use without bundle", func(t *testing.T) {
		ruleErr := goval.NewRuleError(goval.NumberRequired)

		err := tr.Translate(ctx, ruleErr)
		if err != ErrBundleIsNoSet {
			t.Errorf("expect error %v; got %v", ErrBundleIsNoSet, err)
		}
	})
}

func BenchmarkTranslator_Translate(b *testing.B) {
	ctx := context.Background()
	bundle, _ := DefaultBundle()

	tr := NewTranslator(WithBundle(bundle))
	err := &goval.RuleError{
		Code: goval.NumberMin,
		Args: []any{1},
	}
	for i := 0; i < b.N; i++ {
		_ = tr.Translate(ctx, err)
	}
}
