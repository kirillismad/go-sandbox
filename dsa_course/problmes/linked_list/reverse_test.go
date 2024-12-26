package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: NewList([]int{1, 2, 3}),
			want: NewList([]int{3, 2, 1}),
		},
		{
			head: NewList([]int{1, 2}),
			want: NewList([]int{2, 1}),
		},
		{
			head: NewList([]int{1}),
			want: NewList([]int{1}),
		},
		{
			head: nil,
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := reverseList(tt.head)
			require.Equal(t, tt.want, res)
		})

	}
}
