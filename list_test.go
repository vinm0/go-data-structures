package main

import (
	"fmt"
	"testing"
)

func TestPop(t *testing.T) {
	testCases := []struct {
		l       list
		wantPop interface{}
		wantLen int
	}{
		{*NewList(), nil, 0},
		{*sampleList(nil), nil, 0},
		{*sampleList(""), "", 0},
		{*sampleList(0, "", nil), nil, 2},
		{*sampleList(0, 1, 2), 2, 2},
		{*sampleList(0, "two"), "two", 1},
	}

	for _, tC := range testCases {
		testname := fmt.Sprintln(tC.l, ":", tC.wantPop)
		t.Run(testname, func(t *testing.T) {
			pop := tC.l.Pop()
			if len(tC.l.sli) != tC.wantLen {
				t.Errorf("got len %d, want len %d", len(tC.l.sli), tC.wantLen)
			}
			if pop != tC.wantPop {
				t.Errorf("got pop %v, want %v", pop, tC.wantPop)
			}
		})
	}
}

func sampleList(args ...interface{}) *list {
	l := NewList()
	for _, v := range args {
		l.Append(v)
	}

	return l
}
