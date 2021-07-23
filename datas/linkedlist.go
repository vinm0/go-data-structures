package datas

import valid "github.com/vinm0/dataStructures/helper"

type linkedlist struct {
	head *node
	tail *node
	size int
}

type node struct {
	value interface{}
	next  *node
}

func NewLinkedlist() *linkedlist {
	return &linkedlist{}
}

func (ll *linkedlist) Size() int {
	return ll.size
}

func (ll *linkedlist) Append(val interface{}) (ok bool) {
	if val == nil {
		return false
	}

	n := &node{value: val}

	// if linkedlist is empty
	if ll.IsEmpty() {
		return ll.addToEmpty(n)
	}

	ll.tail.next = n
	ll.tail = n
	ll.size++

	return true
}

func (ll *linkedlist) Preppend(val interface{}) (ok bool) {
	n := &node{value: val}

	if ll.IsEmpty() {
		return ll.addToEmpty(n)
	}

	n.next = ll.head.next
	ll.head = n
	ll.size++

	return true
}

func (ll *linkedlist) Insert(val interface{}, index int) (ok bool) {
	if index == 0 {
		return ll.Preppend(val)
	}

	if index == ll.size {
		return ll.Append(val)
	}

	if !valid.ValidIndex(index, ll.size) {
		return false
	}

	n := &node{value: val}

	curr, _ := ll.traverseListTo(index - 1)

	n.next = curr.next
	curr.next = n

	return true
}

func (ll *linkedlist) Index(index int) (val interface{}, ok bool) {
	if ll.IsEmpty() || !valid.ValidIndex(index, ll.size) {
		return nil, false
	}

	if index == ll.size-1 {
		return ll.tail.value, true
	}

	curr, _ := ll.traverseListTo(index)

	return curr.value, true
}

func (ll *linkedlist) Find(val interface{}) (index int, ok bool) {
	if ll.IsEmpty() || !valid.ValidIndex(index, ll.size) {
		return -1, false
	}

	curr := ll.head

	for i := 0; i < ll.size; i++ {
		if curr.value == val {
			return i, true
		}

		curr = curr.next
	}

	return -1, false
}

func (ll *linkedlist) UpdateAt(index int, val interface{}) (ok bool) {
	if ll.IsEmpty() || !valid.ValidIndex(index, ll.size) {
		return false
	}

	if index == ll.size-1 {
		ll.tail.value = val
		return true
	}

	curr, _ := ll.traverseListTo(index)
	curr.value = val

	return true
}

func (ll *linkedlist) RemoveAt(index int) (val interface{}, ok bool) {
	if !valid.ValidIndex(index, ll.size) || ll.size == 0 {
		return nil, false
	}

	if ll.size == 1 {
		val = ll.head.value
		ll.Clear()
		return val, true
	}

	curr, prev := ll.traverseListTo(index)

	prev.next = nil
	ll.tail = prev

	val = curr.value
	curr = nil

	ll.size--

	return val, true
}

func (ll *linkedlist) RemoveVal(val interface{}) (index int, ok bool) {
	if ll.IsEmpty() {
		return -1, false
	}

	if ll.head.value == val {
		ll.head = ll.head.next
		ll.size--
		return 0, true
	}

	if ll.size == 1 {
		return -1, false
	}

	prev := ll.head
	curr := ll.head.next

	for i := 1; i < ll.size; i++ {
		if curr.value == val {
			prev.next = curr.next
			curr = nil
			ll.size--
			return i, true
		}

		prev = curr
		curr = curr.next
	}

	return -1, false
}

func (ll *linkedlist) Clear() {
	ll = new(linkedlist)
}

func (ll *linkedlist) IsEmpty() bool {
	return ll.head == nil
}

func (ll *linkedlist) addToEmpty(n *node) (ok bool) {
	if n == nil {
		return false
	}

	ll.head = n
	ll.tail = n
	ll.size++
	return true
}

func (ll *linkedlist) traverseListTo(index int) (curr *node, prev *node) {

	curr = ll.head

	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}

	return curr, prev
}
