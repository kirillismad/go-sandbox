package linkedlist

// https://leetcode.com/problems/reorder-list/

func reorderList(head *ListNode) {
	slow := head
	fast := head

	// find previous node of middle node
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// first node of second half
	head2 := slow.Next

	// cut the first half
	slow.Next = nil

	// reverse the second half
	var prev *ListNode
	for head2 != nil {
		temp := head2.Next
		head2.Next = prev
		prev = head2
		head2 = temp
	}

	// reorder
	for prev != nil {
		temp := head.Next
		head.Next = prev
		head = prev
		prev = temp
	}
}
