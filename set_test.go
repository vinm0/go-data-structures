package main

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := Set()
	if len(s.data) != 0 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 0)
	}
}

func TestAdd(t *testing.T) {
	s := Set()
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

func TestRemove(t *testing.T) {
	tests := []struct {
		s       set
		rm      interface{}
		wantLen int
	}{
		{*sampleSet(), nil, 3},
		{*sampleSet(), "", 3},
		{*sampleSet(), 0, 3},
		{*sampleSet(), 3, 2},
		{*sampleSet(), "two", 2},
		{*sampleSet(), 1, 2},
		{*Set(), nil, 0},
		{*Set(), "", 0},
		{*Set(), 0, 0},
		{*Set(), 3, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintln(tt.s, ":", tt.wantLen)
		t.Run(testname, func(t *testing.T) {
			tt.s.Remove(tt.rm)
			len := len(tt.s.data)
			if len != tt.wantLen {
				t.Errorf("got len %d, want len %d", len, tt.wantLen)
			}
			if tt.s.Contains(tt.rm) {
				t.Errorf("%v not removed", tt.rm)
			}
		})
	}
}

func TestContains(t *testing.T) {
	samp := *sampleSet()
	tests := []struct {
		s    set
		add  interface{}
		find interface{}
		want bool
	}{
		{samp, nil, nil, false},
		{samp, nil, "", false},
		{samp, "", "", true},
		{samp, nil, 0, false},
		{samp, 0, 0, true},
		{samp, 1, 1, true},
		{samp, 1, "two", true},
		{samp, 1, 3, true},
		{*Set(), nil, nil, false},
		{*Set(), nil, "", false},
		{*Set(), "", "", true},
		{*Set(), nil, 0, false},
		{*Set(), 0, 0, true},
		{*Set(), 1, 1, true},
		{*Set(), 1, "two", false},
		{*Set(), 1, 3, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.s, ": add", tt.add)
		t.Run(testname, func(t *testing.T) {
			tt.s.Add(tt.add)
			if tt.s.Contains(tt.find) != tt.want {
				t.Errorf("got %t, wanted %t", !tt.want, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		s    set
		want []interface{}
	}{
		{*sampleSet(), []interface{}{1, "two", 3}},
		{*Set(), []interface{}{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.want)
		t.Run(testname, func(t *testing.T) {
			sli := tt.s.ToSlice()
			if len(sli) != len(tt.want) {
				t.Errorf("got len %d, wanted len %d", len(sli), len(tt.want))
			}

			for i, v := range sli {
				if v != tt.want[i] {
					t.Errorf("got %v, wanted %v", v, tt.want[i])
				}
			}
		})
	}
}

func sampleSet() *set {
	s := Set()
	s.Add(1)
	s.Add("two")
	s.Add(3)

	return s
}
