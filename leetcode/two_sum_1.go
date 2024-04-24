package leetcode

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	valueIndex := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if index, ok := valueIndex[target-nums[i]]; ok {
			return []int{i, index}
		} else {
			valueIndex[nums[i]] = i
		}
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		if j := search(nums, target-nums[i]); j != -1 {
// 			return []int{i, j}
// 		}
// 	}
// 	return nil
// }
