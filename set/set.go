package set

// Add, Remove, Contains, Size,
type set struct {
	data map[interface{}]struct{}
	size int
}

func (s *set) Add(items ...interface{}) {
	for _, item := range items {
		s.data[item] = struct{}{}
	}
}

func (s *set) Remove(item interface{}) bool {
	_, ok := s.data[item]
	if !ok {
		return false
	}

	delete(s.data, item)
	return true
}
