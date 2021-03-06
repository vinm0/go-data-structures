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
		{List(), nil, 0},
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

func TestUpdateAt(t *testing.T) {
	tests := []struct {
		l        list
		index    int
		val      interface{}
		want     bool
		wantList list
	}{
		{list{}, 1, 1, false, list{}},
		{list{}, 0, nil, false, list{}},
		{list{}, 1, 1, false, list{}},
		{list{nil}, 0, nil, true, list{nil}},
		{list{nil}, 0, 1, true, list{1}},
		{list{nil}, 1, 0, false, list{nil}},
		{list{nil}, -1, 1, false, list{nil}},
		{list{nil, 1}, 1, "two", true, list{nil, "two"}},
		{list{nil, 1}, 0, "NotNil", true, list{"NotNil", 1}},
		{list{1, "two", 3, "four"}, 2, "three", true, list{1, "two", "three", "four"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v : %v", tt.l, tt.val)
		t.Run(testname, func(t *testing.T) {
			ans := tt.l.UpdateAt(tt.index, tt.val)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
			for i, v := range tt.wantList {
				if tt.l[i] != v {
					t.Errorf("got %v at %d, want %v", tt.l[i], i, v)
				}
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	tests := []struct {
		l        list
		index    int
		val      interface{}
		want     bool
		wantList list
	}{
		{list{}, 0, nil, true, list{nil}},
		{list{}, 1, nil, false, list{}},
		{list{}, -1, nil, false, list{}},
		{list{nil}, 0, nil, true, list{nil, nil}},
		{list{nil}, 1, nil, true, list{nil, nil}},
		{list{1, 2}, -1, "three", false, list{1, 2}},
		{list{1, 2}, 3, "three", false, list{1, 2}},
		{list{1, 2}, 2, "three", true, list{1, 2, "three"}},
		{list{1, 2, "three"}, 2, "two", true, list{1, 2, "two", "three"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v : %d,%v", tt.l, tt.index, tt.val)
		t.Run(testname, func(t *testing.T) {
			ans := tt.l.InsertAt(tt.index, tt.val)
			fmt.Println(tt.l)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
			for i, v := range tt.wantList {
				if tt.l[i] != v {
					t.Errorf("got %v at %d, want %v", tt.l[i], i, v)
				}
			}
		})
	}
}

func TestPushFront(t *testing.T) {
	tests := []struct {
		l        list
		val      interface{}
		wantList list
		want     bool
	}{
		{list{}, nil, list{nil}, true},
		{list{}, 0, list{0}, true},
		{list{}, "", list{""}, true},
		{list{nil}, nil, list{nil, nil}, true},
		{list{nil}, 1, list{1, nil}, true},
		{list{0, nil}, "zero", list{"zero", 0, nil}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v : %v", tt.l, tt.val)
		t.Run(testname, func(t *testing.T) {
			ans := tt.l.PushFront(tt.val)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
			for i, v := range tt.wantList {
				if tt.l[i] != v {
					t.Errorf("got %v at %d, want %v", tt.l[i], i, v)
				}
			}
		})
	}
}

// Also satisfies tests for Dequeue
func TestPopFront(t *testing.T) {
	tests := []struct {
		l        list
		wantList list
		want     interface{}
	}{
		{list{}, list{}, nil},
		{list{nil}, list{}, nil},
		{list{0}, list{}, 0},
		{list{""}, list{}, ""},
		{list{"", nil}, list{nil}, ""},
		{list{nil, nil}, list{nil}, nil},
		{list{"test", nil}, list{nil}, "test"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v : %v", tt.l, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := tt.l.PopFront()
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
			for i, v := range tt.wantList {
				if tt.l[i] != v {
					t.Errorf("got %v at %d, want %v", tt.l[i], i, v)
				}
			}
		})
	}
}

// Also satisfies tests for Push and Enqueue
func TestAppend(t *testing.T) {
	tests := []struct {
		l    list
		val  interface{}
		want list
	}{
		{list{}, nil, list{nil}},
		{list{nil}, nil, list{nil, nil}},
		{list{nil}, 0, list{nil, 0}},
		{list{nil}, "", list{nil, ""}},
		{list{1, "two", 3}, "four", list{1, "two", 3, "four"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v : %v", tt.l, tt.val)
		t.Run(testname, func(t *testing.T) {
			tt.l.Append(tt.val)
			if len(tt.l) != len(tt.want) {
				t.Errorf("got len %d, want len %d", len(tt.l), len(tt.want))
			}
			for i, v := range tt.want {
				if i >= len(tt.l) {
					break
				}
				if tt.l[i] != v {
					t.Errorf("got %v at %d, want %v", tt.l[i], i, v)
				}
			}
			if tt.l[len(tt.l)-1] != tt.val {
				t.Errorf("got %v, want %v", tt.l[len(tt.l)-1], tt.val)
			}
		})
	}
}

func sampleList(args ...interface{}) list {
	l := List()
	for _, v := range args {
		l.Append(v)
	}

	return l
}
