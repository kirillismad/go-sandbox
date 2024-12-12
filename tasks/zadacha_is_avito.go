package tasks

import (
	"fmt"
)

func max(numbers []int, left int, right int) int {
	if len(numbers) == 0 {
		panic("empty slice")
	}

	maxIndex := left
	maxValue := numbers[maxIndex]

	for i := left; i <= right; i++ {
		if numbers[i] > maxValue {
			maxIndex = i
			maxValue = numbers[i]
		}
	}
	return maxIndex
}

// https://youtu.be/ou5hSWC82To?si=J7GBkWpoNS2XDhai&t=647
// {1, 3, 1, 2}
// left = 0, right = 3
func solve1(slice []int) int {
	left := 0
	right := len(slice) - 1
	s := 0
	for left <= right {
		maxIndex := max(slice, left, right)
		s = s + (maxIndex-left+1)*slice[maxIndex]
		left = maxIndex + 1
	}
	return s
}

func DemoSolve1() {
	slice := []int{1, 3, 1, 2}
	s := solve1(slice)
	fmt.Println(s)
}
