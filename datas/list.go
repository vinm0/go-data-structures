// The list data structure is an ordered, mutable collection of elements
// with constant time lookup, update, and append operations.
// This GO list uses an underlying slice with methods for manipualation.

package datas

import (
	valid "github.com/vinm0/dataStructures/helper"
)

type list struct {
	sli []interface{}
}

// Create a base empty list
func NewList() *list {
	l := new(list)
	l.sli = make([]interface{}, 0)
	return l
}

// Add item to the end of the list.
func (l *list) Append(val interface{}) {
	l.sli = append(l.sli, val)
}

// Remove item from end of the list.
// Returns value of item removed.
// Returns nil if the list is empty.
func (l *list) Pop() (val interface{}) {
	if l.IsEmpty() {
		return nil
	}

	last := len(l.sli) - 1

	val = l.sli[last]
	l.sli = l.sli[:last]

	return val
}

// Returns the value of the item at the target index.
// Returns nil if target index is out of range.
func (l *list) Index(index int) (val interface{}) {
	if l.IsEmpty() || !valid.ValidIndex(index, len(l.sli)) {
		return nil
	}

	return l.sli[index]
}

// Returns the index of the first occurrence of the target value.
// Returns -1 if the value does not exist.
func (l *list) Find(val interface{}) (index int) {
	if l.IsEmpty() || !valid.ValidIndex(index, len(l.sli)) {
		return -1
	}

	for i, v := range l.sli {
		if v == val {
			return i
		}
	}

	return -1
}

// Sets the item at the target index with the value val.
// Returns true if the list was successfully updated.
// Otherwise, returns false.
func (l *list) UpdateAt(index int, val interface{}) (ok bool) {
	if l.IsEmpty() || !valid.ValidIndex(index, len(l.sli)) {
		return false
	}

	l.sli[index] = val

	return true
}

// Inserts val at the target index.
// The item formerly at index is shifted back.
// Returns true if the list was successfully updated.
// Otherwise, returns false.
func (l *list) InsertAt(index int, val interface{}) (ok bool) {
	if index == len(l.sli) {
		l.Append(val)
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

// Inserts val at index 0.
// All items are shifted 1 index higher.
// Returns true if the list was successfully updated.
// Otherwise, returns false.
func (l *list) Front(val interface{}) (ok bool) {
	var first []interface{}
	first = append(first, val)

	l.sli = append(first, l.sli...)
	return true
}

// Removes the value at index 0.
// All items are shifted 1 index lower.
// Returns the value of the removed item.
// Returns nil if list is empty.
func (l *list) PopFront() (val interface{}) {
	if l.IsEmpty() {
		return nil
	}

	val = l.sli[0]
	l.sli = l.sli[1:]

	return val
}

// Empties the list
func (l *list) Clear() {
	l.sli = make([]interface{}, 0)
}

// Returns true if the list is empty.
// Otherwise, returns false.
func (l *list) IsEmpty() bool {
	return len(l.sli) == 0
}
