package sort

import (
	"cmp"
)

type SelectSorter[T cmp.Ordered] struct{}

func NewSelectSorter[T cmp.Ordered]() SelectSorter[T] {
	return SelectSorter[T]{}
}

func (i SelectSorter[T]) Sort(items []T) {
	for i := 0; i < len(items); i++ {
		minIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}
		items[i], items[minIndex] = items[minIndex], items[i]
	}
}
