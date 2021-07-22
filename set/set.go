package set

// The set data structure implements an unordered collection of elements
// with constant time insertion, lookup, and deletion.
// The GO set uses an underlying map with empty values (i.e, struct{ } ).
type set struct {
	data map[interface{}]struct{}
}

// Creates a base empty set
func NewSet() *set {
	s := &set{}
	s.data = make(map[interface{}]struct{})
	return s
}

func (s *set) Size() int {
	return len(s.data)
}

func (s *set) Add(items ...interface{}) {
	for _, item := range items {
		s.data[item] = struct{}{}
	}
}

func (s *set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.data, item)
	}
}

func (s *set) Contains(item interface{}) bool {
	_, ok := s.data[item]
	return ok
}

func (s *set) ContainsAll(items ...interface{}) bool {
	for _, item := range items {

		_, ok := s.data[item]
		if !ok {
			return false
		}
	}

	return true
}

func (s *set) ContainsAny(items ...interface{}) bool {
	for _, item := range items {

		_, ok := s.data[item]
		if ok {
			return true
		}
	}

	return false
}

func (s *set) Clear() {
	*s = set{}
}

func (s *set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *set) ToSlice() []interface{} {
	setSlice := make([]interface{}, 0, len(s.data))
	for key := range s.data {
		setSlice = append(setSlice, key)
	}

	return setSlice
}
