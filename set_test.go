package main

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet()
	if len(s.data) != 0 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 0)
	}
}

func TestAdd(t *testing.T) {
	s := NewSet()
	if len(s.data) != 0 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 0)
	}

	s.Add("first")
	if len(s.data) != 1 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 1)
	}

	s.Add("second")
	if len(s.data) != 2 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 2)
	}

	s.Remove("second")
	s.Add("2nd")
	if len(s.data) != 2 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 2)
	}

	m := map[int]int{
		1: 2,
		3: 4,
		5: 6,
	}
	s.Add(m)
	if len(s.data) != 2 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 2)
	}

	sli := []string{"1", "2", "3"}
	s.Add(sli)
	if len(s.data) != 2 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 2)
	}

	f := func() int { return 0 }
	s.Add(f)
	if len(s.data) != 2 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 2)
	}

	s.Add(nil)
	if len(s.data) != 2 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 2)
	}

	arr := [5]string{"1", "2", "3", "4", "5"}
	s.Add(arr)
	if len(s.data) != 3 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 3)
	}

	s.Clear()
	s.Add("ten")
	if len(s.data) != 1 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 1)
	}
}
