package queue

import (
	"sandbox/gof"
	"sync"
)

type node[T any] struct {
	value T
	next  *node[T]
}

type ListQueue[T any] struct {
	head *node[T]
	tail *node[T]
	cond *sync.Cond
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{
		cond: sync.NewCond(&sync.Mutex{}),
	}

}

func (q *ListQueue[T]) Enque(item T) error {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	n := &node[T]{value: item}

	if q.head == nil && q.tail == nil {
		q.head = n
		q.tail = n
		return nil
	}

	q.tail.next = n
	q.tail = n
	q.cond.Signal()
	return nil
}

func (q *ListQueue[T]) Deque() (T, error) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for q.head == nil && q.tail == nil {
		q.cond.Wait()
	}

	item := q.head.value

	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return item, nil
}

func (q *ListQueue[T]) Iterator() gof.Iterator[T] {
	return &queue2Iterator[T]{
		q:       q,
		current: q.head,
	}
}

type queue2Iterator[T any] struct {
	q       *ListQueue[T]
	current *node[T]
}

func (i *queue2Iterator[T]) HasNext() bool {
	return i.current != nil
}

func (i *queue2Iterator[T]) Next() (T, error) {
	var item T
	if !i.HasNext() {
		return item, gof.ErrIteratorNext
	}
	item = i.current.value
	i.current = i.current.next
	return item, nil
}
