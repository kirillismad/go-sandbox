package linkedlist

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	left := lists[:len(lists)/2]
	right := lists[len(lists)/2:]

	return mergeTwoLists(mergeKLists(left), mergeKLists(right))
}
