package concurrent

import "sync"

// Set is a concurrent set, safe for read and write operations.
// Guarded by a read-write mutex.
type Set[K comparable] struct {
	mu sync.RWMutex   // Read-write mutex
	m  map[K]struct{} // underlying map
}

// NewSet creates a new concurrent set.
func NewSet[K comparable]() *Set[K] {
	return &Set[K]{
		m: make(map[K]struct{}),
	}
}

// Add adds the given element to the set.
func (s *Set[K]) Add(key K) {
	s.mu.Lock()
	s.m[key] = struct{}{}
	s.mu.Unlock()
}

// Remove removes the given element from the set.
func (s *Set[K]) Remove(key K) {
	s.mu.Lock()
	delete(s.m, key)
	s.mu.Unlock()
}

// Contains checks if the set contains the given element.
func (s *Set[K]) Contains(key K) bool {
	s.mu.RLock()
	_, ok := s.m[key]
	s.mu.RUnlock()
	return ok
}

// Len returns the number of elements in the set.
func (s *Set[K]) Len() int {
	s.mu.RLock()
	length := len(s.m)
	s.mu.RUnlock()
	return length
}

// Values returns all elements in the set.
func (s *Set[K]) Values() []K {
	s.mu.RLock()
	defer s.mu.RUnlock()

	values := make([]K, 0, len(s.m))
	for k := range s.m {
		values = append(values, k)
	}
	return values
}

// Union returns a new set with all elements from both sets.
func (s *Set[K]) Union(other *Set[K]) *Set[K] {
	result := NewSet[K]()
	s.mu.RLock()
	other.mu.RLock()

	defer s.mu.RUnlock()
	defer other.mu.RUnlock()
	for k := range s.m {
		result.Add(k)
	}
	for k := range other.m {
		result.Add(k)
	}

	return result
}

// Intersection returns a new set with elements that are in both sets.
func (s *Set[K]) Intersection(other *Set[K]) *Set[K] {
	result := NewSet[K]()
	s.mu.RLock()
	other.mu.RLock()

	defer s.mu.RUnlock()
	defer other.mu.RUnlock()

	for k := range s.m {
		if other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// Difference returns a new set with elements that are in the first set but not in the second set.
func (s *Set[K]) Difference(other *Set[K]) *Set[K] {
	result := NewSet[K]()
	s.mu.RLock()
	other.mu.RLock()

	defer s.mu.RUnlock()
	defer other.mu.RUnlock()

	for k := range s.m {
		if !other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// SymmetricDifference returns a new set with elements that are in one of the sets but not in both.
func (s *Set[K]) SymmetricDifference(other *Set[K]) *Set[K] {
	result := NewSet[K]()
	s.mu.RLock()
	other.mu.RLock()

	defer s.mu.RUnlock()
	defer other.mu.RUnlock()

	for k := range s.m {
		if !other.Contains(k) {
			result.Add(k)
		}
	}
	for k := range other.m {
		if !s.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// IsSubset checks if the set is a subset of the other set.
func (s *Set[K]) IsSubset(other *Set[K]) bool {
	s.mu.RLock()
	other.mu.RLock()

	defer s.mu.RUnlock()
	defer other.mu.RUnlock()

	for k := range s.m {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the set is a superset of the other set.
func (s *Set[K]) IsSuperset(other *Set[K]) bool {
	return other.IsSubset(s)
}

// Equal checks if the set is equal to the other set.
func (s *Set[K]) Equal(other *Set[K]) bool {
	return s.IsSubset(other) && s.IsSuperset(other)
}

// Clear removes all elements from the set.
func (s *Set[K]) Clear() {
	s.mu.Lock()
	clear(s.m)
	s.mu.Unlock()
}

// Clone returns a new set with a copy of all elements.
func (s *Set[K]) Clone() *Set[K] {
	result := NewSet[K]()
	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.m {
		result.Add(k)
	}
	return result
}
