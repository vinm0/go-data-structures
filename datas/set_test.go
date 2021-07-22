package datas

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet()
	if len(s.data) != 0 {
		t.Errorf("Wrong size set: %d. Expected %d", len(s.data), 0)
	}
}
