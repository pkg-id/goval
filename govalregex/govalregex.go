package govalregex

import "regexp"

const (
	alphaNumericRegexString = "^[a-zA-Z0-9]+$"
)

var (
	AlphaNumeric = regexp.MustCompile(alphaNumericRegexString)
)
