// The set data structure is an unordered collection of elements
// with constant time insertion, lookup, and deletion operations.
// This GO set uses an underlying map with empty values (i.e, struct{ } ).
// The keys of the map serve as the values of the set.
// Currently, does not support non-hashable types
package main

// Uses map to implement set
type set struct {
	data map[interface{}]struct{}
}

// Creates a base empty set
func Set() *set {
	s := &set{}
	s.data = make(map[interface{}]struct{})
	return s
}

// Returns the size of the set
func (s *set) Size() int {
	// Use the length field of the underlying map structure
	return len(s.data)
}

// Adds one item to the set.
// Returns true if item is succesfully added to the set. Otherwise, returns false
func (s *set) Add(item interface{}) bool {

	if !Hashable(item) {
		return false
	}

	// map argument key to empty value
	s.data[item] = struct{}{}
	return true
}

// Removes one item from the set.
// If the item does not exist, no-op.
func (s *set) Remove(item interface{}) {
	if item == nil {
		return
	}

	delete(s.data, item)
}

// Returns true if item exists in the set.
// Otherwise, returns false
func (s *set) Contains(item interface{}) bool {

	if !Hashable(item) {
		return false
	}

	_, ok := s.data[item]
	return ok
}

// Empties the set
func (s *set) Clear() {
	// Value of new set at original address
	s.data = make(map[interface{}]struct{})
}

// Returns true if the set is empty. Otherwise, returns false
func (s *set) IsEmpty() bool {
	return len(s.data) == 0
}

// Returns a slice containing each value in the set.
func (s *set) ToSlice() []interface{} {
	setSlice := make([]interface{}, 0, len(s.data))

	for key := range s.data {
		setSlice = append(setSlice, key)
	}

	return setSlice
}
