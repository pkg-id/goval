package errtrans

import (
	"context"
	"github.com/pkg-id/goval"
	"testing"
)

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
