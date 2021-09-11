package main

import (
	"fmt"
	"testing"
)

func TestHashable(t *testing.T) {
	type testStruct struct{}

	testCases := []struct {
		val  interface{}
		want bool
	}{
		{nil, false},
		{"", true},
		{0, true},
		{[]int{}, false},
		{map[int]int{}, false},
		{func() {}, false},
		{testStruct{}, true},
	}
	for _, tC := range testCases {
		testname := fmt.Sprintf("%T", tC.val)
		t.Run(testname, func(t *testing.T) {
			ans := Hashable(tC.val)
			if ans != tC.want {
				t.Errorf("got %t, want %t", ans, tC.want)
			}
		})
	}
}

func TestValidIndex(t *testing.T) {
	testCases := []struct {
		sli   []interface{}
		index int
		want  bool
	}{
		{[]interface{}{}, 0, false},
		{[]interface{}{}, -1, false},
		{[]interface{}{}, 1, false},
		{[]interface{}{1}, 0, true},
		{[]interface{}{1}, 1, false},
		{[]interface{}{1, "two"}, 1, true},
		{[]interface{}{1, nil}, 1, true},
		{[]interface{}{1, "two"}, 2, false},
		{[]interface{}{1, nil}, 2, false},
	}
	for _, tC := range testCases {
		testname := fmt.Sprintf("%v : %d", tC.sli, tC.index)
		t.Run(testname, func(t *testing.T) {
			ans := ValidIndex(tC.index, len(tC.sli))
			if ans != tC.want {
				t.Errorf("got %t, want %t", ans, tC.want)
			}
		})
	}
}
