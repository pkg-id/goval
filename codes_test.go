package goval

import (
	"testing"
)

type customTypeButSameBase ruleCode

func (c customTypeButSameBase) Equal(other RuleCoder) bool {
	v, ok := other.(customTypeButSameBase)
	return ok && v == c
}

type customTypeButSameRootBase int

func (c customTypeButSameRootBase) Equal(other RuleCoder) bool {
	v, ok := other.(customTypeButSameRootBase)
	return ok && v == c
}

func TestRuleCode_Equal(t *testing.T) {
	tests := []struct {
		desc string
		got  bool
		exp  bool
	}{
		{
			desc: "type and value is same",
			got:  NumberRequired.Equal(NumberRequired),
			exp:  true,
		},
		{
			desc: "type same but value is different",
			got:  NumberRequired.Equal(NumberRequired),
			exp:  true,
		},
		{
			desc: "type and value is same. But, it used value literal",
			got:  NumberRequired.Equal(ruleCode(3000)),
			exp:  true,
		},
		{
			desc: "type is different and literal value is same, but the source type is same",
			got:  NumberRequired.Equal(customTypeButSameBase(3000)),
			exp:  false,
		},
		{
			desc: "type is different and literal value is same, but the root source type is same",
			got:  NumberRequired.Equal(customTypeButSameRootBase(3000)),
			exp:  false,
		},
	}

	for _, tc := range tests {
		if tc.got != tc.exp {
			t.Fatalf("got: %-6v. exp: %-6v. %s", tc.got, tc.exp, tc.desc)
		}
	}
}
