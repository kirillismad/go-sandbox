package leetcode

// https://leetcode.com/problems/contains-duplicate/

// #hashmap

func containsDuplicate(nums []int) bool {
	valueIsContains := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if valueIsContains[nums[i]] {
			return true
		}
		valueIsContains[nums[i]] = true
	}
	return false
}
