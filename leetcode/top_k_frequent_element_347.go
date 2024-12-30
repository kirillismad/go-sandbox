package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)

	// O(n)
	for _, val := range nums {
		counter[val]++
	}

	// O(n)
	keys := make([]int, 0, len(counter))
	for k := range counter {
		keys = append(keys, k)
	}
	// O(n*log(n))
	sort.Slice(keys, func(i, j int) bool { return counter[keys[i]] > counter[keys[j]] })

	return keys[:min(k, len(keys))]
}
