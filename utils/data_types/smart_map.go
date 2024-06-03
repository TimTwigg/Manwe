package data_type_utils

import (
	"fmt"
	"sync"
	"time"
)

type SmartMap[K comparable, V any] struct {
	values      map[K]V
	mu          sync.RWMutex
	LastUpdated time.Time
	Constructor func(K) V
}

func (m *SmartMap[K, V]) Get(key K) V {
	// Lock mutex to prevent multithreaded access to internal reference map.
	m.mu.Lock()
	defer m.mu.Unlock()

	// Initialize internal reference map if it does not already exist.
	if m.values == nil {
		m.values = make(map[K]V)
	}

	// Get value from internal reference map. If it doesn't exist and we have a constructor, use it to create the value.
	value, ok := m.values[key]
	if !ok && m.Constructor != nil {
		if m.Constructor == nil {
			panic(fmt.Sprintf("Key '%v' does not exist in lockable map, and no constructor exists!", key))
		}
		value = m.Constructor(key)
		m.values[key] = value
	}

	// Update the last updated time and return the value.
	m.LastUpdated = time.Now()
	return value
}

func (m *SmartMap[K, V]) Set(key K, value V) {
	// Lock mutex to prevent multithreaded access to internal reference map.
	m.mu.Lock()
	defer m.mu.Unlock()

	// Initialize internal reference map if it does not already exist.
	if m.values == nil {
		m.values = make(map[K]V)
	}

	// Update the last updated time.
	m.LastUpdated = time.Now()

	// Update value.
	m.values[key] = value
}

func (m *SmartMap[K, V]) Has(key K) bool {
	// If internal reference map has not been initialized, key does not exist.
	if m.values == nil {
		return false
	}

	// Lock mutex to prevent multithreaded access to internal reference map.
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if key exists in internal reference map and update the last updated time.
	_, ok := m.values[key]
	m.LastUpdated = time.Now()
	return ok
}

func (m *SmartMap[K, V]) Iter() []KeyValuePair[K, V] {
	// Early return if we've never bothered to initialize the internal reference map.
	if m.values == nil {
		return make([]KeyValuePair[K, V], 0)
	}

	// Lock mutex to prevent multithreaded access to internal reference map.
	m.mu.Lock()
	defer m.mu.Unlock()

	// Convert internal reference map to a slice of key-value pairs.
	pairs := make([]KeyValuePair[K, V], 0, len(m.values))
	for k, v := range m.values {
		pairs = append(pairs, KeyValuePair[K, V]{k, v})
	}

	// Update the last updated time and return the slice of key-value pairs.
	m.LastUpdated = time.Now()
	return pairs
}

func (m *SmartMap[K, V]) Delete(key K) {
	// Early return if we've never bothered to initialize the internal reference map.
	if m.values == nil {
		return
	}

	// Lock mutex to prevent multithreaded access to internal reference map.
	m.mu.Lock()
	defer m.mu.Unlock()

	// Update the last updated time and delete the key from the internal reference map.
	m.LastUpdated = time.Now()
	delete(m.values, key)
}

func (m *SmartMap[K, V]) Keys() []K {
	// Early return if we've never bothered to initialize the internal reference map.
	if m.values == nil {
		return make([]K, 0)
	}

	// Lock mutex to prevent multithreaded access to internal reference map.
	m.mu.Lock()
	defer m.mu.Unlock()

	// Convert internal reference map to a slice of keys.
	keys := make([]K, 0, len(m.values))
	for k := range m.values {
		keys = append(keys, k)
	}

	// Update the last updated time and return the slice of keys.
	m.LastUpdated = time.Now()
	return keys
}

func (m *SmartMap[K, V]) Values() []V {
	// Early return if we've never bothered to initialize the internal reference map.
	if m.values == nil {
		return make([]V, 0)
	}

	// Lock mutex to prevent multithreaded access to internal reference map.
	m.mu.Lock()
	defer m.mu.Unlock()

	// Convert internal reference map to a slice of values.
	values := make([]V, 0, len(m.values))
	for _, v := range m.values {
		values = append(values, v)
	}

	// Update the last updated time and return the slice of values.
	m.LastUpdated = time.Now()
	return values
}
