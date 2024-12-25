package linkedlist

// https://leetcode.com/problems/remove-nth-node-from-end-of-list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	rightNode := dummy
	for i := 0; i < n+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode := dummy

	for rightNode != nil {
		leftNode = leftNode.Next
		rightNode = rightNode.Next
	}

	leftNode.Next = leftNode.Next.Next

	return dummy.Next
}
