package cache

import (
	"github.com/samber/lo"
	"hash/fnv"
)

type FIFOShardsCache[T any] struct {
	shards    []Cache[T]
	numShards int
}

func NewFIFOShardsCache[T any](numShards int, shardSize int) *FIFOShardsCache[T] {
	return &FIFOShardsCache[T]{
		shards:    lo.Times(numShards, func(_ int) Cache[T] { return NewFIFOCache[T](shardSize) }),
		numShards: numShards,
	}
}

func (c FIFOShardsCache[T]) Get(key string) (T, bool) {
	h := fnv.New32()
	_, err := h.Write([]byte(key))
	if err != nil {
		panic(err)
	}
	shard := c.shards[h.Sum32()%uint32(c.numShards)]
	return shard.Get(key)
}

func (c FIFOShardsCache[T]) Put(key string, value T) {
	h := fnv.New32()
	_, err := h.Write([]byte(key))
	if err != nil {
		panic(err)
	}
	shard := c.shards[h.Sum32()%uint32(c.numShards)]
	shard.Put(key, value)
}
