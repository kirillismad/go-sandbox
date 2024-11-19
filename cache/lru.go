package cache

import (
	"container/list"
	"fmt"
)

// LRUCache is a cache based on LRU
type LRUCache struct {
	capacity int                      // Maximum cache capacity
	cache    map[string]*list.Element // Storing data as a hash
	order    *list.List               // Doubly linked list for usage order
}

// Cache entry
type entry struct {
	key   string
	value string
}

// NewLRUCache Creates a new LRU cache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element, capacity),
		order:    list.New(),
	}
}

// Get Gets a value from the cache
func (c *LRUCache) Get(key string) (string, bool) {
	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem) // Move to front as the most recently used
		return elem.Value.(*entry).value, true
	}
	return "", false // Element not found
}

// Put Puts a value into the cache
func (c *LRUCache) Put(key, value string) {
	if elem, found := c.cache[key]; found {
		// Update existing element
		c.order.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	// Add new element
	if c.order.Len() >= c.capacity {
		// If cache is full, remove the last element (least recently used)
		oldest := c.order.Back()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.cache, oldest.Value.(*entry).key)
		}
	}

	// Add new element to the front
	newEntry := &entry{key, value}
	elem := c.order.PushFront(newEntry)
	c.cache[key] = elem
}

func main() {
	cache := NewLRUCache(3)
	cache.Put("A", "value1")
	cache.Put("B", "value2")
	cache.Put("C", "value3")

	fmt.Println(cache.Get("A")) // Output: value1, true

	cache.Put("D", "value4") // Evicts "B" as it is the least recently used

	_, found := cache.Get("B") // Output: false, as "B" was evicted
	fmt.Println("B found:", found)

	fmt.Println(cache.Get("C")) // Output: value3, true
}
