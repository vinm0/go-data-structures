package main

import (
	"fmt"
	"testing"
)

func TestAppendLL(t *testing.T) {
	tests := []struct {
		ll       linkedlist
		val      interface{}
		want     bool
		wantList linkedlist
	}{
		{
			linkedlist{},
			nil,
			true,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			nil,
			true,
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{nil, nil}, size: 2},
		},
		{
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{nil, nil}, size: 2},
			"three",
			true,
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{"three", nil}, size: 3},
		},
		{
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{nil, nil}, size: 2},
			0,
			true,
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{0, nil}, size: 3},
		},
		{
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{nil, nil}, size: 2},
			"",
			true,
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{"", nil}, size: 3},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List Append test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.Append(tt.val)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
			if tt.ll.head.value != tt.wantList.head.value {
				t.Errorf("got %v head, want %v head", *tt.ll.head, *tt.wantList.head)
			}
			if tt.ll.tail.value != tt.wantList.tail.value {
				t.Errorf("got %v tail, want %v tail", *tt.ll.tail, *tt.wantList.tail)
			}
			if tt.ll.size != tt.wantList.size {
				t.Errorf("got %v size, want %v size", tt.ll.size, tt.wantList.size)
			}
		})
	}
}

func TestPushFrontLL(t *testing.T) {
	tests := []struct {
		ll       linkedlist
		val      interface{}
		want     bool
		wantList linkedlist
	}{
		{
			linkedlist{},
			nil,
			true,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			nil,
			true,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 2},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 2},
			"three",
			true,
			linkedlist{head: &node{"three", nil}, tail: &node{nil, nil}, size: 3},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 2},
			0,
			true,
			linkedlist{head: &node{0, nil}, tail: &node{nil, nil}, size: 3},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 2},
			"",
			true,
			linkedlist{head: &node{"", nil}, tail: &node{nil, nil}, size: 3},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List PushFront test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.PushFront(tt.val)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
			if tt.ll.head.value != tt.wantList.head.value {
				t.Errorf("got %v head, want %v head", *tt.ll.head, *tt.wantList.head)
			}
			if tt.ll.tail.value != tt.wantList.tail.value {
				t.Errorf("got %v tail, want %v tail", *tt.ll.tail, *tt.wantList.tail)
			}
			if tt.ll.size != tt.wantList.size {
				t.Errorf("got %v size, want %v size", tt.ll.size, tt.wantList.size)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		ll       linkedlist
		index    int
		val      interface{}
		want     bool
		wantList linkedlist
	}{
		{
			linkedlist{},
			0,
			nil,
			true,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			0,
			nil,
			true,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 2},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 2},
			1,
			"three",
			true,
			linkedlist{head: &node{nil, &node{"three", &node{nil, nil}}}, tail: &node{nil, nil}, size: 3},
		},
		{
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{nil, nil}, size: 2},
			-1,
			"three",
			false,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 2},
		},
		{
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{nil, nil}, size: 2},
			2,
			"three",
			true,
			linkedlist{head: &node{nil, &node{nil, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
		},
		{
			linkedlist{head: &node{nil, &node{nil, nil}}, tail: &node{nil, nil}, size: 2},
			0,
			"three",
			true,
			linkedlist{head: &node{"three", &node{nil, &node{nil, nil}}}, tail: &node{nil, nil}, size: 3},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List PushFront test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.Insert(tt.val, tt.index)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
			if tt.ll.size != tt.wantList.size {
				t.Errorf("got %v size, want %v size", tt.ll.size, tt.wantList.size)
			}
			for i := 0; i < tt.ll.size; i++ {
				v1 := tt.ll.Index(i)
				v2 := tt.wantList.Index(i)
				if v1 != v2 {
					t.Errorf("got %v at %d, want %v", v1, i, v2)
				}
			}
		})
	}
}

func TestIndexLL(t *testing.T) {
	tests := []struct {
		ll    linkedlist
		index int
		want  interface{}
	}{
		{linkedlist{}, 0, nil},
		{linkedlist{}, -1, nil},
		{linkedlist{}, 1, nil},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			0,
			nil,
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			0,
			"one",
		},
		{
			linkedlist{head: &node{nil, &node{nil, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			2,
			"three",
		},
		{
			linkedlist{head: &node{nil, &node{nil, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			1,
			nil,
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			1,
			5,
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			-1,
			nil,
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			3,
			nil,
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List Index test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.Index(tt.index)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestFindLL(t *testing.T) {
	tests := []struct {
		ll   linkedlist
		find interface{}
		want int
	}{
		{linkedlist{}, nil, -1},
		{linkedlist{}, "", -1},
		{linkedlist{}, 0, -1},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			nil,
			0,
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			0,
			-1,
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			"one",
			0,
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			nil,
			-1,
		},
		{
			linkedlist{head: &node{nil, &node{nil, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			"three",
			2,
		},
		{
			linkedlist{head: &node{nil, &node{nil, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			nil,
			0,
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			5,
			1,
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			"five",
			-1,
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List Find test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.Find(tt.find)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestUpdateAtLL(t *testing.T) {
	tests := []struct {
		ll       linkedlist
		index    int
		val      interface{}
		want     bool
		wantList linkedlist
	}{
		{linkedlist{}, 0, nil, false, linkedlist{}},
		{linkedlist{}, 0, "", false, linkedlist{}},
		{linkedlist{}, 0, 0, false, linkedlist{}},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			0,
			"",
			true,
			linkedlist{head: &node{"", nil}, tail: &node{"", nil}, size: 1},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			0,
			nil,
			true,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			-1,
			"two",
			false,
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			1,
			"two",
			false,
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			0,
			"two",
			true,
			linkedlist{head: &node{"two", nil}, tail: &node{"two", nil}, size: 1},
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			1,
			"two",
			true,
			linkedlist{head: &node{nil, &node{"two", &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List UpdateAt test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.UpdateAt(tt.index, tt.val)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
			if tt.ll.size != tt.wantList.size {
				t.Errorf("got %v size, want %v size", tt.ll.size, tt.wantList.size)
			}
			for i := 0; i < tt.ll.size; i++ {
				v1 := tt.ll.Index(i)
				v2 := tt.wantList.Index(i)
				if v1 != v2 {
					t.Errorf("got %v at %d, want %v", v1, i, v2)
				}
			}
		})
	}
}

func TestRemoveAtLL(t *testing.T) {
	tests := []struct {
		ll       linkedlist
		index    int
		want     interface{}
		wantList linkedlist
	}{
		{linkedlist{}, 0, nil, linkedlist{}},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			0,
			nil,
			linkedlist{},
		},
		{
			linkedlist{head: &node{0, nil}, tail: &node{0, nil}, size: 1},
			0,
			0,
			linkedlist{},
		},
		{
			linkedlist{head: &node{"", nil}, tail: &node{"", nil}, size: 1},
			0,
			"",
			linkedlist{},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			-1,
			nil,
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			1,
			nil,
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			0,
			"one",
			linkedlist{},
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			1,
			5,
			linkedlist{head: &node{nil, &node{"three", nil}}, tail: &node{"three", nil}, size: 2},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List RemoveAt test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.RemoveAt(tt.index)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
			if tt.ll.size != tt.wantList.size {
				t.Errorf("got %v size, want %v size", tt.ll.size, tt.wantList.size)
			}
			if tt.ll.head != nil && tt.wantList.head != nil &&
				tt.ll.head.value != tt.wantList.head.value {
				t.Errorf("got %v head, want %v head", *tt.ll.head, *tt.wantList.head)
			}
			if tt.ll.tail != nil && tt.wantList.tail != nil &&
				tt.ll.tail.value != tt.wantList.tail.value {
				t.Errorf("got %v tail, want %v tail", *tt.ll.tail, *tt.wantList.tail)
			}
			for i := 0; i < tt.ll.size; i++ {
				v1 := tt.ll.Index(i)
				v2 := tt.wantList.Index(i)
				if v1 != v2 {
					t.Errorf("got %v at %d, want %v", v1, i, v2)
				}
			}
		})
	}
}

func TestRemoveValLL(t *testing.T) {
	tests := []struct {
		ll       linkedlist
		val      interface{}
		want     int
		wantList linkedlist
	}{
		{linkedlist{}, nil, -1, linkedlist{}},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			nil,
			0,
			linkedlist{},
		},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			"",
			-1,
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
		},
		{
			linkedlist{head: &node{"", nil}, tail: &node{"", nil}, size: 1},
			nil,
			-1,
			linkedlist{head: &node{"", nil}, tail: &node{"", nil}, size: 1},
		},
		{
			linkedlist{head: &node{0, nil}, tail: &node{0, nil}, size: 1},
			0,
			0,
			linkedlist{},
		},
		{
			linkedlist{head: &node{"", nil}, tail: &node{"", nil}, size: 1},
			"",
			0,
			linkedlist{},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			nil,
			-1,
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			1,
			-1,
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
		},
		{
			linkedlist{head: &node{"one", nil}, tail: &node{"one", nil}, size: 1},
			"one",
			0,
			linkedlist{},
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			5,
			1,
			linkedlist{head: &node{nil, &node{"three", nil}}, tail: &node{"three", nil}, size: 2},
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List RemoveVal test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans := tt.ll.RemoveVal(tt.val)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
			if tt.ll.size != tt.wantList.size {
				t.Errorf("got %v size, want %v size", tt.ll.size, tt.wantList.size)
			}
			if tt.ll.head != nil && tt.wantList.head != nil &&
				tt.ll.head.value != tt.wantList.head.value {
				t.Errorf("got %v head, want %v head", *tt.ll.head, *tt.wantList.head)
			}
			if tt.ll.tail != nil && tt.wantList.tail != nil &&
				tt.ll.tail.value != tt.wantList.tail.value {
				t.Errorf("got %v tail, want %v tail", *tt.ll.tail, *tt.wantList.tail)
			}
			for i := 0; i < tt.ll.size; i++ {
				v1 := tt.ll.Index(i)
				v2 := tt.wantList.Index(i)
				if v1 != v2 {
					t.Errorf("got %v at %d, want %v", v1, i, v2)
				}
			}
		})
	}
}

func TestClearLL(t *testing.T) {
	emptyList := linkedlist{}
	tests := []struct {
		ll   linkedlist
		want *linkedlist
	}{
		{linkedlist{}, &emptyList},
		{
			linkedlist{head: &node{nil, nil}, tail: &node{nil, nil}, size: 1},
			&emptyList,
		},
		{
			linkedlist{head: &node{"", nil}, tail: &node{"", nil}, size: 1},
			&emptyList,
		},
		{
			linkedlist{head: &node{0, nil}, tail: &node{0, nil}, size: 1},
			&emptyList,
		},
		{
			linkedlist{head: &node{nil, &node{5, &node{"three", nil}}}, tail: &node{"three", nil}, size: 3},
			&emptyList,
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Linked List RemoveVal test %d", i)
		t.Run(testname, func(t *testing.T) {
			tt.ll.Clear()
			if tt.ll.size != tt.want.size {
				t.Errorf("got %v size, want %v size", tt.ll.size, tt.want.size)
			}
			if tt.ll.head != nil || tt.ll.tail != nil {
				t.Errorf("Linked List contains non-nil node fields")
			}
		})
	}
}
