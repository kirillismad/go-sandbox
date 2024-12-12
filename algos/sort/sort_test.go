package sort

import (
	"math/rand"
	gosort "sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSorter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		sorter Sorter[int]
	}{
		{
			name:   "Select",
			sorter: NewSelectSorter[int](),
		},
		{
			name:   "Insert",
			sorter: NewInsertSorter[int](),
		},
		{
			name:   "Shell",
			sorter: NewShellSorter[int](),
		},
		{
			name:   "Bubble",
			sorter: NewBubbleSorter[int](),
		},
		{
			name:   "Quick",
			sorter: NewQuickSorter[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			items := rand.Perm(1 << 8)

			rand.Shuffle(len(items), func(i, j int) {
				items[i], items[j] = items[j], items[i]
			})

			tt.sorter.Sort(items)

			require.True(t, gosort.IsSorted(gosort.IntSlice(items)))
		})
	}
}

func BenchmarkSorter(b *testing.B) {
	tests := []struct {
		name   string
		sorter Sorter[int]
	}{
		{
			name:   "Select",
			sorter: NewSelectSorter[int](),
		},
		{
			name:   "Insert",
			sorter: NewInsertSorter[int](),
		},
		{
			name:   "Shell",
			sorter: NewShellSorter[int](),
		},
		{
			name:   "Bubble",
			sorter: NewBubbleSorter[int](),
		},
		{
			name:   "Quick",
			sorter: NewQuickSorter[int](),
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				items := rand.Perm(1 << 8)

				b.StartTimer()
				tt.sorter.Sort(items)
			}
		})
	}
}
