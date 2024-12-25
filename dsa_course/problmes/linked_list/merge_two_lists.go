package linkedlist

// https://leetcode.com/problems/merge-two-sorted-lists

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	dummy := new(ListNode)
	// Переменная curr указывает на текущий узел результирующего списка (изначально dummy).
	// Цикл продолжается до тех пор, пока хотя бы один из списков a или b не станет пустым.
	for curr := dummy; a != nil || b != nil; curr = curr.Next {
		// Если текущий узел из списка a существует (a != nil) и либо:
		// список b закончился (b == nil),
		// либо значение текущего узла из a меньше (a.Val < b.Val),
		// тогда:
		// Указатель curr.Next привязывается к текущему узлу из a.
		// Указатель a перемещается на следующий узел (a = a.Next).
		if a != nil && (b == nil || a.Val < b.Val) {
			curr.Next = a
			a = a.Next
		} else {
			// Иначе (если список b не пуст и его текущий элемент меньше либо равен текущему элементу из a):
			// Указатель curr.Next привязывается к текущему узлу из b.
			// Указатель b перемещается на следующий узел (b = b.Next).
			curr.Next = b
			b = b.Next
		}
	}
	return dummy.Next
}
