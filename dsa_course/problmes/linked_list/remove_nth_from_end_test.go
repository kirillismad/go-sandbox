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
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			n:    2,
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			n:    1,
			want: &ListNode{Val: 1},
		},

		{
			head: &ListNode{Val: 1},
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
