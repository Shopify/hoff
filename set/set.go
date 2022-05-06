package set

// Set implements a unique, unordered set of homogeneous element types.
// An empty struct is used for the value, to signify that only its presence matters, not its value.
type Set[T comparable] map[T]struct{}

// FromSlice creates and returns a new Set using the supplied values and their type.
func FromSlice[T comparable](values []T) Set[T] {
	set := make(Set[T], len(values))
	set.Add(values...)
	return set
}

// Add inserts a value to the set, without retaining insertion order.
func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

// Delete removes all matching elements from the set.
func (s Set[T]) Delete(values ...T) {
	for _, value := range values {
		delete(s, value)
	}
}

// Has returns a boolean, for whether the value exists in the Set.
func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}

// Len returns the number of items in the Set.
func (s Set[T]) Len() int {
	return len(s)
}

// ForEach accepts and applies a callback function to each element in the Set.
func (s Set[T]) ForEach(it func(T)) {
	for value := range s {
		it(value)
	}
}

// Values returns all elements of the set, unordered.
func (s Set[T]) Values() []T {
	values := make([]T, 0, s.Len())
	s.ForEach(
		func(value T) {
			values = append(values, value)
		},
	)
	return values
}

// Clone returns a new copy of the original Set.
func (s Set[T]) Clone() Set[T] {
	set := make(Set[T])
	set.Add(s.Values()...)
	return set
}

// Union returns a Set of the combined Sets.
// Example: [1, 2].Union([2, 3]) = [1, 2, 3].
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := s.Clone()
	union.Add(other.Values()...)
	return union
}

// Intersection returns a Set containing all the common items between both Sets.
// Example: [1, 2, 3].Intersection([2, 3, 4]) = [2, 3].
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := make(Set[T])
	s.ForEach(
		func(value T) {
			if other.Has(value) {
				intersection.Add(value)
			}
		},
	)
	return intersection
}

// Difference returns a Set containing all the items not contained in the other set.
// Note this is "unidirectional", and the result is _only_ the elements in A that are not in B.
// Example: [1, 2, 3].Difference([2, 3, 4]) = [1].
func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := make(Set[T])
	s.ForEach(
		func(value T) {
			if !other.Has(value) {
				diff.Add(value)
			}
		},
	)
	return diff
}

// SymmetricalDifference returns the inverse of Intersection, a Set containing all elements
// not common to both incoming Sets.
// Example: [1, 2, 3].SymmetricalDifference([2, 3, 4]) = [1, 4].
func (s Set[T]) SymmetricalDifference(other Set[T]) Set[T] {
	// [1, 2, 3].Intersection([2, 3, 4]) = [2, 3]
	intersection := s.Intersection(other)
	// [1, 2, 3].Union([2, 3, 4]) = [1, 2, 3, 4]
	union := s.Union(other)
	// [1, 2, 3, 4].Difference([2, 3]) = [1, 4]
	symmetricalDifference := union.Difference(intersection)
	return symmetricalDifference
}

// Equals whether the sets have exactly the same elements.
func (s Set[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for value := range s {
		if !other.Has(value) {
			return false
		}
	}
	return true
}
