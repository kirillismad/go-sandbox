package leetcode

// https://leetcode.com/problems/contains-duplicate-ii/

// #hashmap

func containsNearbyDuplicate(nums []int, k int) bool {
	valueIndex := map[int]int{}

	for i, value := range nums {
		if nearestIndex, ok := valueIndex[value]; ok {
			if i-nearestIndex <= k {
				return true
			}
		}
		valueIndex[value] = i
	}
	return false
}
