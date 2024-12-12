package stack

import (
	"container/list"
	"sandbox/gof"
)

type InfStack3[T any] struct {
	linkedList *list.List
}

func NewInfStack3[T any]() *InfStack3[T] {
	return &InfStack3[T]{
		linkedList: list.New(),
	}
}

func (s *InfStack3[T]) Push(item T) error {
	s.linkedList.PushFront(item)
	return nil
}

func (s *InfStack3[T]) Pop() (T, error) {
	head := s.linkedList.Front()
	if head == nil {
		var item T
		return item, ErrStackIsEmpty
	}

	item := head.Value.(T)
	s.linkedList.Remove(head)
	return item, nil
}

func (s *InfStack3[T]) Iterator() gof.Iterator[T] {
	return &infStackIterator3[T]{
		stack:   s,
		current: s.linkedList.Front(),
	}
}

type infStackIterator3[T any] struct {
	stack   *InfStack3[T]
	current *list.Element
}

func (i *infStackIterator3[T]) Next() (T, error) {
	var item T
	if !i.HasNext() {
		return item, gof.ErrIteratorNext
	}
	item = i.current.Value.(T)

	i.current = i.current.Next()

	return item, nil
}

func (i *infStackIterator3[T]) HasNext() bool {
	return i.current != nil
}
