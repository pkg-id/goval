package goval

import (
	"encoding/json"
	"fmt"
)

type RuleCode int

const (
	ruleCodeNilBase RuleCode = (1 + iota) * 1_000
	ruleCodeBaseString
	ruleCodeBaseNumber
	ruleCodeBaseSlice
	ruleCodeBaseMap
)

const (
	NilRequired = ruleCodeNilBase + iota
)

const (
	StringRequired = ruleCodeBaseString + iota
	StringMin
	StringMax
	StringMatch
)

const (
	NumberRequired = ruleCodeBaseNumber + iota
	NumberMin
	NumberMax
)

const (
	SliceRequired = ruleCodeBaseSlice + iota
	SliceMin
	SliceMax
	SliceEach
)

const (
	MapRequired = ruleCodeBaseMap + iota
	MapMin
	MapMax
	MapEach
)

func IsCodeEqual(a, b RuleCode) bool {
	return a == b
}

type RuleError struct {
	Code  RuleCode `json:"code"`
	Input any      `json:"input"`
	Args  []any    `json:"args"`
}

type auxRuleError RuleError

func NewRuleError(code RuleCode, input any, args ...any) *RuleError {
	return &RuleError{
		Code:  code,
		Input: input,
		Args:  args,
	}
}

func (r *RuleError) Error() string {
	return r.String()
}

func (r *RuleError) String() string {
	b, err := r.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func (r *RuleError) MarshalJSON() ([]byte, error) {
	aux := auxRuleError(*r)
	return json.Marshal(aux)
}

type RuleErrors struct {
	Code RuleCode
	Errs []RuleError
	Args []any
}

func NewRuleErrors(code RuleCode, errs []RuleError, args ...any) *RuleErrors {
	return &RuleErrors{
		Code: code,
		Errs: errs,
		Args: args,
	}
}

func (e *RuleErrors) Error() string { return e.String() }

func (e *RuleErrors) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		return "goval: RuleErrors.String: " + err.Error()
	}
	return string(b)
}

func (e *RuleErrors) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(e.Errs)
	if err != nil {
		return nil, fmt.Errorf("goval: RuleErrors.MarshalJSON: %w", err)
	}

	return b, nil
}
