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
		{NewList(), nil, 0},
		{sampleList(nil), nil, 0},
		{sampleList(""), "", 0},
		{sampleList(0, "", nil), nil, 2},
		{sampleList(0, 1, 2), 2, 2},
		{sampleList(0, "two"), "two", 1},
	}

	for _, tC := range testCases {
		testname := fmt.Sprintln(tC.l, ":", tC.wantPop)
		t.Run(testname, func(t *testing.T) {
			pop := tC.l.Pop()
			if len(tC.l) != tC.wantLen {
				t.Errorf("got len %d, want len %d", len(tC.l), tC.wantLen)
			}
			if pop != tC.wantPop {
				t.Errorf("got pop %v, want %v", pop, tC.wantPop)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		l     list
		index int
		want  interface{}
	}{
		{list{}, 0, nil},
		{list{nil}, 0, nil},
		{list{1}, 0, 1},
		{list{1}, -1, nil},
		{list{1}, 1, nil},
		{list{1, "two"}, 1, "two"},
		{list{1, "two"}, 3, nil},
	}
	for _, tC := range testCases {
		testname := fmt.Sprintf("%v : %d", tC.l, tC.index)
		t.Run(testname, func(t *testing.T) {
			ans := tC.l.Index(tC.index)
			if ans != tC.want {
				t.Errorf("got %v, want %v", ans, tC.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		l    list
		find interface{}
		want int
	}{
		{list{}, nil, -1},
		{list{}, 0, -1},
		{list{nil}, nil, 0},
		{list{0, nil}, nil, 1},
		{list{0, nil}, 0, 0},
		{list{0, nil}, 1, -1},
		{list{0, "one"}, "one", 1},
		{list{0, "one"}, "two", -1},
		{list{"one", "one"}, "one", 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v : %d", tt.l, tt.find)
		t.Run(testname, func(t *testing.T) {
			ans := tt.l.Find(tt.find)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func sampleList(args ...interface{}) list {
	l := NewList()
	for _, v := range args {
		l.Append(v)
	}

	return l
}
