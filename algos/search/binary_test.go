package search

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	slice := []int{-1, 0, 3, 5, 9, 12}
	target := 12

	index := binarySearch(slice, target)
	fmt.Printf("binarySearch(%v, %v) = %v\n", slice, target, index)
}

func Test_binarySearchR(t *testing.T) {
	slice := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 5

	index := binarySearchR(slice, target, 0, len(slice))
	fmt.Printf("binarySearchR(%v, %v, %v, %v) = %v\n", slice, target, 0, len(slice), index)
}
