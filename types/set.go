package types

// Set stores unique values of a comparable type.
type Set[T comparable] map[T]struct{}

// Add inserts value into the set.
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Has reports whether value belongs to the set.
func (s Set[T]) Has(value T) bool {
	_, exists := s[value]
	return exists
}

// Remove deletes value from the set.
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Size returns the number of values in the set.
func (s Set[T]) Size() int {
	return len(s)
}

// ToSlice returns the set values in unspecified order.
func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for value := range s {
		slice = append(slice, value)
	}
	return slice
}

// NewSet returns an empty set.
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}
