package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reorderList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: NewList([]int{1, 2, 3, 4}),
			want: NewList([]int{1, 4, 2, 3}),
		},
		{
			head: NewList([]int{1, 2, 3}),
			want: NewList([]int{1, 3, 2}),
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			reorderList(tt.head)

			require.Equal(t, tt.want, tt.head)
		})
	}
}
