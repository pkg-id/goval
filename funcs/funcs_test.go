package funcs_test

import (
	"github.com/pkg-id/goval/funcs"
	"reflect"
	"sort"
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

func TestMap(t *testing.T) {
	outs := funcs.Map([]int{1, 2, 3}, func(inp int) int {
		return inp * 2
	})

	exp := []int{2, 4, 6}
	if !reflect.DeepEqual(outs, exp) {
		t.Fatalf("expect %v; got %v", exp, outs)
	}
}

func TestValues(t *testing.T) {
	inp := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	outs := funcs.Values(inp)
	sort.Ints(outs)
	exp := []int{1, 2, 3}
	if !reflect.DeepEqual(outs, exp) {
		t.Fatalf("expect %v; got %v", exp, outs)
	}
}

func TestReduce(t *testing.T) {
	sum := funcs.Reduce([]int{1, 2, 3}, 0, func(acc, inp int) int {
		return acc + inp
	})

	exp := 1 + 2 + 3
	if sum != exp {
		t.Fatalf("expect %v; got %v", exp, sum)
	}
}
