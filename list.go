// The list data structure is an ordered, mutable collection of elements
// with constant time lookup, update, and append operations.
// This GO list uses an underlying slice with methods for manipualation.

package main

type list []interface{}

// Create a base empty list
func List() list {
	l := make([]interface{}, 0)
	return l
}

// Add item to the end of the list.
func (l *list) Append(val interface{}) {
	*l = append(*l, val)
}

// Remove item from end of the list.
// Returns value of item removed.
// Returns nil if the list is empty.
func (l *list) Pop() (val interface{}) {
	if l.IsEmpty() {
		return nil
	}

	last := len(*l) - 1

	val = (*l)[last]
	*l = (*l)[:last]

	return val
}

// Returns the value of the item at the target index.
// Returns nil if target index is out of range.
func (l *list) Index(index int) (val interface{}) {
	if l.IsEmpty() || !ValidIndex(index, len(*l)) {
		return nil
	}

	return (*l)[index]
}

// Returns the index of the first occurrence of the target value.
// Returns -1 if the value does not exist.
func (l *list) Find(val interface{}) (index int) {
	if l.IsEmpty() || !ValidIndex(index, len(*l)) {
		return -1
	}

	for i, v := range *l {
		if v == val {
			return i
		}
	}

	return -1
}

// Sets the item at the target index with the value val.
// Returns true if the list was successfully updated,
// even if the previous value is identical to the updated value.
// Otherwise, returns false.
func (l *list) UpdateAt(index int, val interface{}) (ok bool) {
	if l.IsEmpty() || !ValidIndex(index, len(*l)) {
		return false
	}

	(*l)[index] = val

	return true
}

// Inserts val at the target index.
// The item formerly at index is shifted back.
// Returns true if the list was successfully updated.
// Otherwise, returns false.
func (l *list) InsertAt(index int, val interface{}) (ok bool) {
	if index == len(*l) {
		l.Append(val)
		return true
	}

	if index == 0 {
		return l.PushFront(val)
	}

	if !ValidIndex(index, len(*l)) {
		return false
	}

	// Copy beginning of slice to prevent direct alteration
	beg := make([]interface{}, index+1)
	copy(beg, (*l)[:index])

	end := (*l)[index:]

	*l = append(beg, end...)
	(*l)[index] = val

	return true
}

// Inserts val at index 0.
// All items are shifted 1 index higher.
// Returns true if the list was successfully updated.
// Otherwise, returns false.
func (l *list) PushFront(val interface{}) (ok bool) {
	var first []interface{}
	first = append(first, val)

	*l = append(first, *l...)
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

	val = (*l)[0]
	*l = (*l)[1:]

	return val
}

func (l *list) Push(val interface{}) {
	l.Append(val)
}

func (l *list) Enqueue(val interface{}) {
	l.Append(val)
}

func (l *list) Dequeue() (val interface{}) {
	return l.PopFront()
}

// Empties the list
func (l *list) Clear() {
	*l = make([]interface{}, 0)
}

// Returns true if the list is empty.
// Otherwise, returns false.
func (l *list) IsEmpty() bool {
	return len(*l) == 0
}
