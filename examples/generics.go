package examples

import "fmt"

func DemoGenerics() {
	demoList()
	demoMapKeys()
}
func demoMapKeys() {
	fmt.Println(newline + fname(demoMapKeys))
	m := map[int]string{
		1: "111",
		3: "333",
		5: "555",
	}
	keys := MapKeys(m)

	fmt.Println(keys)

}
func demoList() {
	fmt.Println(newline + fname(demoList))
	lst := List[int]{}
	// lst := List[string]{}

	s := []int{1, 2, 3, 4, 5}
	// s := []string{"A", "B", "C", "D", "E"}
	for _, i := range s {
		lst.Add(i)
	}

	// listIter := ListIterator[int]{curr: lst.head}
	listIter := lst.GetIterator()
	for listIter.HasNext() {
		next := listIter.Next()
		fmt.Println(next)
	}
}

type Element[T interface{}] struct {
	value T
	next  *Element[T]
}

func (e Element[T]) String() string {
	return fmt.Sprintf("Element(value:%v)", e.value)
}

type List[T interface{}] struct {
	head *Element[T]
}

func (l *List[T]) Add(value T) {
	element := &Element[T]{value: value}
	if l.head == nil {
		l.head = element
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = element
	}
}

func (l *List[T]) GetIterator() *ListIterator[T] {
	return &ListIterator[T]{curr: l.head}
}

type ListIterator[T interface{}] struct {
	curr *Element[T]
}

func (li *ListIterator[T]) Next() *Element[T] {
	ret := li.curr
	li.curr = li.curr.next
	return ret
}

func (li *ListIterator[T]) HasNext() bool {
	return li.curr != nil
}

func MapKeys[T comparable, V any](m map[T]V) []T {
	r := make([]T, 0, len(m))

	for k := range m {
		r = append(r, k)
	}
	return r
}
