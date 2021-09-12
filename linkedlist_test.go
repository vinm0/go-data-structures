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
