package sort

import (
	"cmp"
)

type InsertSorter[T cmp.Ordered] struct{}

func NewInsertSorter[T cmp.Ordered]() InsertSorter[T] {
	return InsertSorter[T]{}
}

func (i InsertSorter[T]) Sort(items []T) {
	for i := 1; i < len(items); i++ {
		for j := i; j > 0 && items[j] < items[j-1]; j-- {
			items[j], items[j-1] = items[j-1], items[j]
		}
	}
}
