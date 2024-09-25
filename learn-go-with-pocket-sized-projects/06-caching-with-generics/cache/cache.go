package cache

import (
	"sync"
	"time"
)

type valueWithTimeout[V any] struct {
	value   V
	expires time.Time // After this time, the value is useless
}

// Cache is key-value storage.
type Cache[K comparable, V any] struct {
	ttl  time.Duration
	mu   sync.Mutex
	data map[K]valueWithTimeout[V]
}

// New create a new Cache with an initialised data.
func New[K comparable, V any](ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:  ttl,
		data: make(map[K]valueWithTimeout[V]),
	}
}

// Read returns the associated value for a key,
// and a boolean to true if the key is absent.
func (c *Cache[K, V]) Read(key K) (V, bool) {
	// Lock the reading and the possible writing on the map
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroV V
	v, found := c.data[key]

	switch {
	case !found:
		return zeroV, false
	case v.expires.Before(time.Now()):
		delete(c.data, key)
		return zeroV, false
	default:
		return v.value, true
	}
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	// Lock the writing on the map
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = valueWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
}

// Delete removes the entry for the given key.
// If the key isn't present, Delete is a no-op.
func (c *Cache[K, V]) Delete(key K) {
	// Lock the deletion on the map
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
