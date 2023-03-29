package govalregex

import (
	"regexp"
	"sync"
)

const (
	alphaNumericRegexString = "^[a-zA-Z0-9]+$"
)

var (
	// AlphaNumeric is a lazy-compiled regex for alphanumeric characters only.
	AlphaNumeric = NewLazy(alphaNumericRegexString)
)

// LazyCompiler is a lazy-compiled regex.
// That is, the regex is compiled only when the RegExp method is called.
type LazyCompiler struct {
	once     sync.Once
	expr     string
	compiled *regexp.Regexp
}

// NewLazy creates a new LazyCompiler.
func NewLazy(expr string) *LazyCompiler {
	return &LazyCompiler{
		once:     sync.Once{},
		expr:     expr,
		compiled: nil,
	}
}

// RegExp returns the compiled regular expression.
func (l *LazyCompiler) RegExp() *regexp.Regexp {
	// Compile the regex only once, and cache it.
	l.once.Do(func() {
		l.compiled = regexp.MustCompile(l.expr)
	})
	return l.compiled
}
