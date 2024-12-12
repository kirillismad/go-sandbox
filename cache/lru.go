package cache

import (
	"container/list"
	"sync"
)

type lruEntity[T any] struct {
	key   string
	value T
}

type LRUCache[T any] struct {
	capacity int
	cache    map[string]*list.Element
	order    *list.List
	mu       sync.Mutex
}

func NewLRUCache[T any](capacity int) Cache[T] {
	return &LRUCache[T]{
		capacity: capacity,
		cache:    make(map[string]*list.Element, capacity),
		order:    list.New(),
	}
}

func (c *LRUCache[T]) Get(key string) (T, bool) {
	var zero T

	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem)
		return elem.Value.(*lruEntity[T]).value, true
	}

	return zero, false
}

func (c *LRUCache[T]) Put(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem)
		elem.Value.(*lruEntity[T]).value = value
		return
	}

	if c.order.Len() >= c.capacity {
		oldest := c.order.Back()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.cache, oldest.Value.(*lruEntity[T]).key)
		}
	}

	newEntry := &lruEntity[T]{key: key, value: value}
	elem := c.order.PushFront(newEntry)
	c.cache[key] = elem
}
