package tasks

import (
	"hash/fnv"
	"sync"
)

func main() {}

// 1. Subscribe
// Logic: Creates a new channel, registers it for the given topic, and returns it to the caller.
// Constraint: To handle bursty traffic, the returned channel must have a buffer size of 10.
// 2. Publish
// Logic: Sends the message to all active subscribers of that topic.
// Constraint (Non-Blocking): If a subscriber’s channel is full (buffer is maxed out),
// the publisher must drop the message for that specific subscriber and continue to the next one immediately.
// The system cannot wait/block for slow consumers.

// sharding to decrease lock contention
type ShardedSyncMap[T any] struct {
	buckets []*SyncMap[string, T]
}

func NewShardedSyncMap[T any](shards int) *ShardedSyncMap[T] {
	maps := make([]*SyncMap[string, T], 0, shards)
	for range shards {
		maps = append(maps, NewSyncMap[string, T]())
	}
	return &ShardedSyncMap[T]{
		buckets: maps,
	}
}

func (m *ShardedSyncMap[T]) hash(k string) uint64 {
	hasher := fnv.New64a()
	_, _ = hasher.Write([]byte(k))
	return hasher.Sum64()
}

func (m *ShardedSyncMap[T]) PutElem(k string, el T) {
	h := int(m.hash(k) % uint64(len(m.buckets)))
	m.buckets[h].PutElem(k, el)
}

func (m *ShardedSyncMap[T]) Get(k string) ([]T, bool) {
	h := int(m.hash(k) % uint64(len(m.buckets)))
	return m.buckets[h].Get(k)
}

type SyncMap[K comparable, T any] struct {
	storage map[K][]T

	mu sync.RWMutex
}

func NewSyncMap[K comparable, T any]() *SyncMap[K, T] {
	return &SyncMap[K, T]{
		storage: make(map[K][]T),
	}
}

func (m *SyncMap[K, T]) PutElem(k K, el T) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.storage[k] = append(m.storage[k], el)
}

func (m *SyncMap[K, T]) Get(k K) ([]T, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	elems, ok := m.storage[k]
	return elems, ok
}

type PubSub struct {
	chans *ShardedSyncMap[chan string]
}

func NewPubSub(shards int) *PubSub {
	return &PubSub{
		chans: NewShardedSyncMap[chan string](shards),
	}
}

func (s *PubSub) Subscribe(topic string) <-chan string {
	ch := make(chan string, 10)
	s.chans.PutElem(topic, ch)

	return ch
}

func (s *PubSub) Publish(topic string, message string) {
	chans, ok := s.chans.Get(topic)
	if !ok {
		return
	}
	for _, ch := range chans {
		select {
		case ch <- message:
		default:
		}
	}
}
