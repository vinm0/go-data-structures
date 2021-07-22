// The set data structure implements an unordered collection of elements
// with constant time insertion, lookup, and deletion.
package set

// The GO set uses an underlying map with empty values (i.e, struct{ } ).
// The keys of the map serve as the values of the set.
type set struct {
	data map[interface{}]struct{}
}

// Creates a base empty set
func NewSet() *set {
	s := &set{}
	s.data = make(map[interface{}]struct{})
	return s
}

// Returns the size of the set
func (s *set) Size() int {
	return len(s.data)
}

// Adds one or more items to the set
// Note: Slices should be unpacked.
func (s *set) Add(items ...interface{}) {
	for _, item := range items {
		s.data[item] = struct{}{}
	}
}

// Removes one or more items from the set
// Note: Slices should be unpacked.
func (s *set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.data, item)
	}
}

// Returns true if item exists in the set.
// Otherwise, returns false
func (s *set) Contains(item interface{}) bool {
	_, ok := s.data[item]
	return ok
}

// Returns true if all items exist in the set.
// Otherwise, returns false.
// Note: Slices should be unpacked.
func (s *set) ContainsAll(items ...interface{}) bool {
	for _, item := range items {

		_, ok := s.data[item]
		if !ok {
			return false
		}
	}

	return true
}

// Returns true if any item provided exists in the set.
// Returns false if no item exists.
// Note: Slices should be unpacked.
func (s *set) ContainsAny(items ...interface{}) bool {
	for _, item := range items {

		_, ok := s.data[item]
		if ok {
			return true
		}
	}

	return false
}

// Empties the set
func (s *set) Clear() {
	*s = set{}
}

// Returns true if the set is empty. Otherwise, returns false
func (s *set) IsEmpty() bool {
	return len(s.data) == 0
}

// Returns a slice containing each value in the set.
// Tip: Minimize calls to ToSlice()
func (s *set) ToSlice() []interface{} {
	setSlice := make([]interface{}, 0, len(s.data))
	for key := range s.data {
		setSlice = append(setSlice, key)
	}

	return setSlice
}
