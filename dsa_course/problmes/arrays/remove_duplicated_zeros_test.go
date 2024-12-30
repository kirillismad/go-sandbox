package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicatedZeros(t *testing.T) {
	t.Parallel()
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{1, 1, 0, 0, 2, 3, 0, 3, 0, 0, 0, 1},
			want: []int{1, 1, 0, 2, 3, 0, 3, 0, 1},
		},
		{
			arr:  []int{1, 0, 0, 1, 1, 1},
			want: []int{1, 0, 1, 1, 1},
		},
		{
			arr:  []int{0},
			want: []int{0},
		},
		{
			arr:  []int{0, 0, 0, 0, 0},
			want: []int{0},
		},
		{
			arr:  []int{0, 1, 0},
			want: []int{0, 1, 0},
		},
		{
			arr:  []int{0, 0, 1, 0, 0},
			want: []int{0, 1, 0},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := RemoveDuplicatedZeros(tt.arr)
			require.ElementsMatch(t, tt.want, res)
		})
	}
}
