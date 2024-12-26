package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_middleNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: NewList([]int{1, 2, 3, 4, 5}),
			want: NewList([]int{3, 4, 5}),
		},
		{
			head: NewList([]int{1, 2, 3, 4, 5, 6}),
			want: NewList([]int{4, 5, 6}),
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := middleNode(tt.head)
			require.Equal(t, tt.want, res)
		})
	}
}
