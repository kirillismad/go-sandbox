package linkedlist

// https://leetcode.com/problems/reverse-linked-list/description/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}
