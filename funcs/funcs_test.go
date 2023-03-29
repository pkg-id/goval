package funcs_test

import (
	"github.com/pkg-id/goval/funcs"
	"testing"
)

func TestContains(t *testing.T) {
	ok := funcs.Contains([]int{1, 2, 4}, func(value int) bool { return value == 2 })
	if !ok {
		t.Fatalf("expect ok")
	}

	ok = funcs.Contains([]int{1, 3, 4}, func(value int) bool { return value == 2 })
	if ok {
		t.Fatalf("expect not ok")
	}
}
