package cache

import (
	"container/list"
	"sync"
)

// LRUCache is a cache based on LRU
type LRUCache[T any] struct {
	capacity int                      // Maximum cache capacity
	cache    map[string]*list.Element // Storing data as a hash
	order    *list.List               // Doubly linked list for usage order
	mu       sync.Mutex               // Mutex for thread safety
}

// NewLRUCache Creates a new LRU cache with the given capacity
func NewLRUCache[T any](capacity int) *LRUCache[T] {
	return &LRUCache[T]{
		capacity: capacity,
		cache:    make(map[string]*list.Element, capacity),
		order:    list.New(),
	}
}

// Get Gets a value from the cache
func (c *LRUCache[T]) Get(key string) (T, bool) {
	var zero T

	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem) // Move to front as the most recently used
		return elem.Value.(*Entry[T]).value, true
	}

	return zero, false // Element not found
}

// Put Puts a value into the cache
func (c *LRUCache[T]) Put(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.cache[key]; found {
		// Update existing element
		c.order.MoveToFront(elem)
		elem.Value.(*Entry[T]).value = value
		return
	}

	// Add new element
	if c.order.Len() >= c.capacity {
		// If cache is full, remove the last element (least recently used)
		oldest := c.order.Back()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.cache, oldest.Value.(*Entry[T]).key)
		}
	}

	// Add new element to the front
	newEntry := &Entry[T]{key: key, value: value}
	elem := c.order.PushFront(newEntry)
	c.cache[key] = elem
}
