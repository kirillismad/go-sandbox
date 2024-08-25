package sort

import "cmp"

type BubbleSorter[T cmp.Ordered] struct {
}

func NewBubbleSorter[T cmp.Ordered]() *BubbleSorter[T] {
	return &BubbleSorter[T]{}
}

func (s *BubbleSorter[T]) Sort(items []T) {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[i] {
				items[j], items[i] = items[i], items[j]
			}
		}
	}
}
