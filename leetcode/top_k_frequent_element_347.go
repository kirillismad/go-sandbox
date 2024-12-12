package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	elementCounter := make(map[int]int)

	// O(n)
	for _, val := range nums {
		elementCounter[val]++
	}

	// O(n)
	keys := make([]int, 0, len(elementCounter))
	for k := range elementCounter {
		keys = append(keys, k)
	}
	// O(n*log(n))
	sort.Slice(keys, func(i, j int) bool { return elementCounter[keys[i]] > elementCounter[keys[j]] })

	// O(k)
	result := make([]int, 0)
	for i := 0; i < k && i < len(keys); i++ {
		result = append(result, keys[i])
	}

	return result
}
