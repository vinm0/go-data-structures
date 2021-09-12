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
