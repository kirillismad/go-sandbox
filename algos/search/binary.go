package search

func binarySearch(nums []int, target int) int {
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

func binarySearchR(slice []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	midIndex := left + (right-left)/2
	if target > slice[midIndex] {
		return binarySearchR(slice, target, midIndex+1, right)
	} else if target < slice[midIndex] {
		return binarySearchR(slice, target, left, right-1)
	} else {
		return midIndex
	}
}
