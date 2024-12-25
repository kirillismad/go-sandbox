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
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		},
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := middleNode(tt.head)
			require.Equal(t, tt.want, res)
		})
	}
}
