package cache

import (
	"sync"
)

type clockEntry[T any] struct {
	key       string
	value     T
	reference bool
}

type ClockCache[T any] struct {
	capacity int
	cache    map[string]*clockEntry[T]
	entries  []*clockEntry[T]
	hand     int
	mu       sync.Mutex
}

func NewClockCache[T any](capacity int) *ClockCache[T] {
	return &ClockCache[T]{
		capacity: capacity,
		cache:    make(map[string]*clockEntry[T], capacity),
		entries:  make([]*clockEntry[T], 0, capacity),
		hand:     0,
	}
}

func (c *ClockCache[T]) Get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, found := c.cache[key]; found {
		entry.reference = true
		return entry.value, true
	}

	var zero T
	return zero, false
}

func (c *ClockCache[T]) Put(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If the key already exists, update the value and set the reference bit to true
	if entry, found := c.cache[key]; found {
		entry.value = value
		entry.reference = true
		return
	}

	// If the cache is not full, add the new entry to the cache
	if len(c.entries) < c.capacity {
		entry := &clockEntry[T]{key: key, value: value, reference: true}
		c.entries = append(c.entries, entry)
		c.cache[key] = entry
		return
	}

	// If the cache is full, find the first entry with reference bit set to false
	// and replace it with the new entry
	for {
		current := c.entries[c.hand]
		if !current.reference {
			delete(c.cache, current.key)
			current.key = key
			current.value = value
			current.reference = true
			c.cache[key] = current
			c.hand = (c.hand + 1) % c.capacity
			return
		}
		current.reference = false
		c.hand = (c.hand + 1) % c.capacity
	}
}
