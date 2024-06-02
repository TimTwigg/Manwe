package utils

import (
	"fmt"
)

// LockableMap is a map that can be locked to prevent further writes.
// The idea is to use this for our internal, default maps so that we can rest assured that they will not be modified.
type LockableMap[K comparable, V any] struct {
	values map[K]V
	locked bool

	Constructor func(K) V
}

func (m *LockableMap[K, V]) Get(key K) V {
	// Initialize internal reference map if it does not already exist.
	if m.values == nil {
		m.values = make(map[K]V)
	}

	// Get value from internal reference map. We will panic if it doesn't exist and we don't have a constructor, or if the struct is locked.
	value, ok := m.values[key]
	if !ok {
		if m.Constructor == nil {
			panic(fmt.Sprintf("Key '%v' does not exist in lockable map, and no constructor exists!", key))
		}
		if m.locked {
			panic(fmt.Sprintf("Cannot create new value in locked map! (Using get on key '%v' that does not exist)", key))
		}

		value = m.Constructor(key)
		m.values[key] = value
	}
	return value
}

func (m *LockableMap[K, V]) Set(key K, value V) {
	// Lock mutex to prevent multithreaded access to internal reference map.
	if m.locked {
		panic("Cannot set value in locked map!")
	}

	// Initialize internal reference map if it does not already exist.
	if m.values == nil {
		m.values = make(map[K]V)
	}

	// Update value.
	m.values[key] = value
}

func (m *LockableMap[K, V]) Has(key K) bool {
	// If internal reference map has not been initialized, key does not exist.
	if m.values == nil {
		return false
	}

	// Check if key exists in internal reference map and update the last updated time.
	_, ok := m.values[key]
	return ok
}

func (m *LockableMap[K, V]) Iter() []KeyValuePair[K, V] {
	// Early return if we've never bothered to initialize the internal reference map.
	if m.values == nil {
		return make([]KeyValuePair[K, V], 0)
	}

	// Convert internal reference map to a slice of key-value pairs.
	pairs := make([]KeyValuePair[K, V], 0, len(m.values))
	for k, v := range m.values {
		pairs = append(pairs, KeyValuePair[K, V]{k, v})
	}

	// Return the slice of key-value pairs.
	return pairs
}

func (m *LockableMap[K, V]) Delete(key K) {
	// Early return if we've never bothered to initialize the internal reference map.
	if m.values == nil {
		return
	}

	// Delete the key from the internal reference map.
	delete(m.values, key)
}

func (m *LockableMap[K, V]) ToString() string {
	starterString := "LockableMap {"
	for key, value := range m.values {
		starterString += fmt.Sprintf("\n\t%v: %v,", key, value)
	}
	starterString += "\n}"
	return starterString
}

func (m *LockableMap[K, V]) Lock() {
	m.locked = true
}
