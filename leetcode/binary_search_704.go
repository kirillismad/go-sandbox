package leetcode

// https://leetcode.com/problems/binary-search/description/
// #binary-search

func search(nums []int, target int) int {
	leftBorder := 0
	rightBorder := len(nums) - 1

	for leftBorder <= rightBorder {
		midIndex := leftBorder + (rightBorder-leftBorder)/2

		switch {
		case target < nums[midIndex]:
			rightBorder = midIndex - 1
		case target > nums[midIndex]:
			leftBorder = midIndex + 1
		default:
			return midIndex
		}
	}
	return -1
}

// func search(nums []int, target int) int {
// 	i, found := sort.Find(len(nums), func(i int) int { return cmp.Compare(target, nums[i]) })

// 	if found {
// 		return i
// 	}
// 	return -1
// }
