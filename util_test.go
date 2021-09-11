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
