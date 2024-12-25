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
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
		},
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}},
		},
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			reorderList(tt.head)

			require.Equal(t, tt.want, tt.head)
		})
	}
}
