package govalregex

import "testing"

func TestLazyCompiler_RegExp(t *testing.T) {
	re := NewLazy(`^[a-z]+$`)
	if re.compiled != nil {
		t.Errorf("expect no value until first request is created")
	}

	c1 := re.RegExp()
	if c1 == nil {
		t.Fatalf("expect not nil")
	}

	c2 := re.RegExp()
	if c1 != c2 {
		t.Errorf("expect c1 and c2 is equal")
	}
}
