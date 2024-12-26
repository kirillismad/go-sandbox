package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeKLists(t *testing.T) {
	t.Parallel()
	tests := []struct {
		lists []*ListNode
		want  *ListNode
	}{
		{
			lists: []*ListNode{
				NewList([]int{1, 4, 5}),
				NewList([]int{1, 3, 4}),
				NewList([]int{2, 6}),
			},
			want: NewList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			lists: []*ListNode{},
			want:  nil,
		},
		{
			lists: []*ListNode{nil},
			want:  nil,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, tt.want, mergeKLists(tt.lists))
		})
	}
}
