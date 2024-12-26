package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeTwoLists(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a    *ListNode
		b    *ListNode
		want *ListNode
	}{
		{
			a:    NewList([]int{1, 2, 4}),
			b:    NewList([]int{1, 3, 4}),
			want: NewList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			a:    NewList([]int{1, 2}),
			b:    NewList([]int{1, 3}),
			want: NewList([]int{1, 1, 2, 3}),
		},
		{
			a:    NewList([]int{1}),
			b:    nil,
			want: NewList([]int{1}),
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := mergeTwoLists(tt.a, tt.b)
			require.Equal(t, tt.want, res)
		})
	}
}
