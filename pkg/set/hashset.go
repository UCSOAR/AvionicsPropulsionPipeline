package set

type HashSet[T comparable] struct {
	inner map[T]struct{}
}

func NewHashSet[T comparable]() HashSet[T] {
	inner := make(map[T]struct{})
	return HashSet[T]{inner: inner}
}

func (s *HashSet[T]) Put(item T) {
	s.inner[item] = struct{}{}
}

func (s *HashSet[T]) Remove(item T) {
	delete(s.inner, item)
}

func (s *HashSet[T]) Has(item T) bool {
	_, exists := s.inner[item]
	return exists
}
