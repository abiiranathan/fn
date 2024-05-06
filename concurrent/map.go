package concurrent

import "sync"

// Map is a concurrent map, safe for read and write operations.
// Guarded by a read-write mutex.
type Map[K comparable, V any] struct {
	mu sync.RWMutex // Read-write mutex
	m  map[K]V      // underlying map
}

// NewMap creates a new concurrent map.
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: make(map[K]V),
	}
}

// Get returns the value associated with the given key.
// If the key is not found, Get returns the zero value for the value type and false.
func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	m.mu.RLock()
	value, ok = m.m[key]
	m.mu.RUnlock()
	return
}

// Set sets the given value to the given key in the map.
func (m *Map[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

// Delete deletes the item with the given key from the map.
func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

// Len returns the number of items in the map.
func (m *Map[K, V]) Len() int {
	m.mu.RLock()
	length := len(m.m)
	m.mu.RUnlock()
	return length
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.mu.RLock()
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
	m.mu.RUnlock()
}

// Keys returns all keys in the map.
func (m *Map[K, V]) Keys() []K {
	m.mu.RLock()
	keys := make([]K, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	m.mu.RUnlock()
	return keys
}

// Values returns all values in the map.
func (m *Map[K, V]) Values() []V {
	m.mu.RLock()
	values := make([]V, 0, len(m.m))
	for _, v := range m.m {
		values = append(values, v)
	}
	m.mu.RUnlock()
	return values
}

// Clear removes all items from the map.
func (m *Map[K, V]) Clear() {
	m.mu.Lock()
	clear(m.m)
	m.mu.Unlock()
}
