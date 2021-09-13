package main

type linkedlist struct {
	head *node
	tail *node
	size int
}

type node struct {
	value interface{}
	next  *node
}

// Creates a base empty linked list.
func Linkedlist() *linkedlist {
	return &linkedlist{}
}

// Returns the number of elements in the linked list
func (ll *linkedlist) Size() int {
	return ll.size
}

// Adds one value to the end of the linked list.
// Returns true if the value is successfully added to the linked list.
// Otherwise, returns false.
func (ll *linkedlist) Append(val interface{}) (ok bool) {

	n := &node{value: val}

	if ll.IsEmpty() {
		return ll.addToEmpty(n)
	}

	// Update tail
	ll.tail.next = n
	ll.tail = n
	ll.size++

	return true
}

// Adds one value to the beginning of the linked list.
// Returns true if the value is successfully added to the linked list.
// Otherwise, returns false.
func (ll *linkedlist) PushFront(val interface{}) (ok bool) {
	n := &node{value: val}

	if ll.IsEmpty() {
		return ll.addToEmpty(n)
	}

	// Update head
	n.next = ll.head.next
	ll.head = n
	ll.size++

	return true
}

// Adds a new value at the target index.
// The previous value at the target index will follow the new index.
// Returns true if the value is successfully added to the linked list.
// Otherwise, returns false.
func (ll *linkedlist) Insert(val interface{}, index int) (ok bool) {
	if index == 0 {
		return ll.PushFront(val)
	}

	if index == ll.size {
		return ll.Append(val)
	}

	if !ValidIndex(index, ll.size) {
		return false
	}

	n := &node{value: val}

	// Traverse to the index prior to the target index
	curr, _ := ll.traverseListTo(index - 1)

	n.next = curr.next
	curr.next = n
	ll.size++

	return true
}

// Returns the value of the item at the target index.
// Returns nil if index is invalid or linked list is empty.
func (ll *linkedlist) Index(index int) (val interface{}) {
	if ll.IsEmpty() || !ValidIndex(index, ll.size) {
		return nil
	}

	if index == ll.size-1 {
		return ll.tail.value
	}

	curr, _ := ll.traverseListTo(index)

	return curr.value
}

// Returns the index if the first instance if the target value.
// Returns -1 if the target value does not exist.
func (ll *linkedlist) Find(val interface{}) (index int) {
	if ll.IsEmpty() {
		return -1
	}

	curr := ll.head

	for i := 0; i < ll.size; i++ {
		if curr.value == val {
			return i
		}

		curr = curr.next
	}

	return -1
}

// Updates the value of the element at the target index.
// Returns true if update is successful. Otherwise, return false.
func (ll *linkedlist) UpdateAt(index int, val interface{}) (ok bool) {
	if ll.IsEmpty() || !ValidIndex(index, ll.size) {
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

// Removes the item at the target index.
// Returns the value of the removed item.
// Returns nil if no item is removed.
func (ll *linkedlist) RemoveAt(index int) (val interface{}) {
	if ll.IsEmpty() || !ValidIndex(index, ll.size) {
		return nil
	}

	if index == 0 {
		val = ll.head.value
		ll.head = ll.head.next
		ll.size--

		return val
	}

	curr, prev := ll.traverseListTo(index)

	prev.next = curr.next

	val = curr.value
	curr = nil

	if index == ll.size-1 {
		ll.tail = prev
	}

	ll.size--

	return val
}

// Removes the first occurrence of the target value.
// Returns the index of the item removed.
// Returns -1 if no item is removed.
func (ll *linkedlist) RemoveVal(val interface{}) (index int) {
	// Handle instances where the linked list has 0 or 1 items
	// to avoid runtime errors with prev and curr pointers.

	// No items to remove in an empty linked list
	if ll.IsEmpty() {
		return -1
	}

	// Target is the first item.
	if ll.head.value == val {
		ll.head = ll.head.next
		ll.size--
		return 0
	}

	if ll.size == 1 {
		return -1
	}

	prev := ll.head
	curr := ll.head.next

	for i := 1; i < ll.size; i++ {
		if curr.value == val {
			prev.next = curr.next
			curr = nil
			ll.size--

			return i // Value found
		}

		prev = curr
		curr = curr.next
	}

	return -1 // Value not found
}

// Empties the linked list
func (ll *linkedlist) Clear() {
	*ll = linkedlist{}
}

// Returns true if the linked list contains no items.
// Otherwise, returns false.
func (ll *linkedlist) IsEmpty() bool {
	return ll.head == nil
}

// Helper Function: Inserts node into an empty linked list.
// Updates head, tail, and size.
// Returns true if n is a valid node and is succesfully added to the linked list.
// Otherwise, returns false
func (ll *linkedlist) addToEmpty(n *node) (ok bool) {
	if n == nil {
		return false
	}

	ll.head = n
	ll.tail = n
	ll.size++

	return true
}

// Helper Function: Traverses the linked list up to the target index.
// Returns curr as the node at the target index, and prev as the node prior.
// This function does not validate the index argument. It assumes index has
// already been validated and the linked list is not empty.
func (ll *linkedlist) traverseListTo(index int) (curr *node, prev *node) {

	curr = ll.head

	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}

	return curr, prev
}
