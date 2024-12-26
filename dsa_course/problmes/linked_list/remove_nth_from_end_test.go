package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeNthFromEnd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		{
			head: NewList([]int{1, 2, 3, 4, 5}),
			n:    2,
			want: NewList([]int{1, 2, 3, 5}),
		},
		{
			head: NewList([]int{1, 2}),
			n:    1,
			want: NewList([]int{1}),
		},
		{
			head: NewList([]int{1}),
			n:    1,
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := removeNthFromEnd(tt.head, tt.n)
			require.Equal(t, tt.want, res)
		})
	}
}
