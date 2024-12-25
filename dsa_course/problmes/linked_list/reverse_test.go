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
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}},
		},
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
		},
		{
			head: &ListNode{Val: 1, Next: nil},
			want: &ListNode{Val: 1, Next: nil},
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
