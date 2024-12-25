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
			a:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
			b:    &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
		},
		{
			a:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			b:    &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}},
		},
		{
			a:    &ListNode{Val: 1, Next: nil},
			b:    nil,
			want: &ListNode{Val: 1, Next: nil},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := mergeTwoLists(tt.a, tt.b)
			require.Equal(t, tt.want, res)
		})
	}
}
