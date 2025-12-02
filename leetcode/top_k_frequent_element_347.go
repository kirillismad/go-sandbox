package leetcode

import (
	"container/heap"
)

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	// O(n)
	for _, val := range nums {
		counter[val]++
	}

	// O(n)
	maxHeap := make(MaxHeap, 0, len(counter))
	for key, value := range counter {
		heap.Push(&maxHeap, Pair{key: key, count: value})
	}

	result := make([]int, 0, k)
	// O(k log n)
	for range min(k, len(counter)) {
		pair := heap.Pop(&maxHeap).(Pair)
		result = append(result, pair.key)
	}
	return result
}

// min heap
type Pair struct {
	key   int
	count int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
