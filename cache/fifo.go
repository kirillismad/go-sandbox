package cache

import (
	"container/list"
	"sync"
)

type fifoEntry[T any] struct {
	key   string
	value T
}

type FIFOCache[T any] struct {
	capacity int
	cache    map[string]*list.Element
	order    *list.List
	mu       sync.RWMutex
}

func NewFIFOCache[T any](capacity int) Cache[T] {
	return &FIFOCache[T]{
		capacity: capacity,
		cache:    make(map[string]*list.Element, capacity),
		order:    list.New(),
	}
}

func (c *FIFOCache[T]) Get(key string) (T, bool) {
	var zero T

	c.mu.RLock()
	defer c.mu.RUnlock()

	if elem, found := c.cache[key]; found {
		return elem.Value.(*fifoEntry[T]).value, true
	}

	return zero, false
}

func (c *FIFOCache[T]) Put(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.cache[key]; found {
		elem.Value.(*fifoEntry[T]).value = value
		return
	}

	if c.order.Len() >= c.capacity {
		oldest := c.order.Front()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.cache, oldest.Value.(*fifoEntry[T]).key)
		}
	}

	newEntry := &fifoEntry[T]{key: key, value: value}
	elem := c.order.PushBack(newEntry)
	c.cache[key] = elem
}
