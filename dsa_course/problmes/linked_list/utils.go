package linkedlist

func NewList(vals []int) *ListNode {
	dummy := &ListNode{}

	cur := dummy
	for _, val := range vals {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return dummy.Next
}
