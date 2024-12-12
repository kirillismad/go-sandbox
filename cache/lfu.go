package cache

import (
	"container/heap"
	"sync"
)

type lfuEntry[T any] struct {
	key       string
	value     T
	frequency int
	index     int
}

type lfuHeap[T any] []*lfuEntry[T]

func (h lfuHeap[T]) Len() int { return len(h) }

func (h lfuHeap[T]) Less(i, j int) bool {
	return h[i].frequency < h[j].frequency
}

func (h lfuHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *lfuHeap[T]) Push(x interface{}) {
	entry := x.(*lfuEntry[T])
	entry.index = len(*h)
	*h = append(*h, entry)
}

func (h *lfuHeap[T]) Pop() interface{} {
	old := *h
	n := len(old)
	entry := old[n-1]
	old[n-1] = nil
	entry.index = -1
	*h = old[0 : n-1]
	return entry
}

type LFUCache[T any] struct {
	capacity int
	cache    map[string]*lfuEntry[T]
	freqHeap *lfuHeap[T]
	mu       sync.Mutex
}

func NewLFUCache[T any](capacity int) Cache[T] {
	h := &lfuHeap[T]{}
	heap.Init(h)
	return &LFUCache[T]{
		capacity: capacity,
		cache:    make(map[string]*lfuEntry[T], capacity),
		freqHeap: h,
	}
}

func (c *LFUCache[T]) Get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, found := c.cache[key]; found {
		entry.frequency++
		heap.Fix(c.freqHeap, entry.index)
		return entry.value, true
	}

	var zero T
	return zero, false
}

func (c *LFUCache[T]) Put(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, found := c.cache[key]; found {
		entry.value = value
		entry.frequency++
		heap.Fix(c.freqHeap, entry.index)
		return
	}

	if len(c.cache) >= c.capacity {
		lfu := heap.Pop(c.freqHeap).(*lfuEntry[T])
		delete(c.cache, lfu.key)
	}

	entry := &lfuEntry[T]{key: key, value: value, frequency: 1}
	heap.Push(c.freqHeap, entry)
	c.cache[key] = entry
}
