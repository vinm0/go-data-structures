package datas

import (
	valid "github.com/vinm0/dataStructures/helper"	"golang.org/x/tools/go/ssa/interp"
)


type list struct {
	sli []interface{}
}

func NewList() *list {
	l := new(list)
	l.sli = make([]interface{}, 0)
	return l
}

func (l *list) Push(val interface{}) {
	l.sli = append(l.sli, val)
}

func (l *list) Pop() (val interface{}) {
	last := len(l.sli) - 1

	val = l.sli[last]
	l.sli = l.sli[:last]

	return val
}

func (l *list) Index(index int) (val interface{}) {
	if l.IsEmpty() || !valid.ValidIndex(index, len(l.sli)) {
		return nil
	}

	return l.sli[index]
}

func (l *list) Find(val interface{}) (index int) {
	if l.IsEmpty() || !valid.ValidIndex(index, len(l.sli)) {
		return -1
	}

	for i := 0; i < len(l.sli); i++ {
		if l.sli[i] == val {
			return i
		}
	}

	return -1
}

func (l *list) UpdateAt(index int, val interface{}) (ok bool) {
	if l.IsEmpty() || !valid.ValidIndex(index, len(l.sli)) {
		return false
	}

	l.sli[index] = val

	return true
}

func (l *list) InsertAt(index int, val interface{}) (ok bool) {
	if index == len(l.sli) {
		l.Push(val)
		return true
	}

	if index == 0 {
		return l.Front(val)
	}

	if !valid.ValidIndex(index, len(l.sli)) {
		return false
	}

	beg := append(l.sli[:index], val)
	end := l.sli[index:]

	l.sli = append(beg, end...)

	return true
}

func (l *list) Front(val interface{}) (ok bool) {
	l.sli = append([]interface{val}, l.sli...)
	return true
}

func (l *list) PopFront() (val interface{}) {
	if l.IsEmpty() {
		return nil
	}

	val = l.sli[0]
	l.sli = l.sli[1:]

	return val
}

func (l *list) Clear() {
	l.sli = make([]interface{}, 0)
}

func (l *list) IsEmpty() bool {
	return len(l.sli) == 0
}
