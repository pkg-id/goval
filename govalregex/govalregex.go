package govalregex

import (
	"regexp"
	"sync"
)

const (
	alphaNumericRegexString = "^[a-zA-Z0-9]+$"
)

var (
	AlphaNumeric = NewLazy(alphaNumericRegexString)
)

type LazyCompiler struct {
	once     sync.Once
	expr     string
	compiled *regexp.Regexp
}

func NewLazy(expr string) *LazyCompiler {
	return &LazyCompiler{
		once:     sync.Once{},
		expr:     expr,
		compiled: nil,
	}
}

func (l *LazyCompiler) RegExp() *regexp.Regexp {
	l.once.Do(func() {
		l.compiled = regexp.MustCompile(l.expr)
	})
	return l.compiled
}
