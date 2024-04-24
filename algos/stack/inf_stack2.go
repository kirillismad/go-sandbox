package stack

import "sandbox/gof"

type node[T any] struct {
	item T
	next *node[T]
}

type InfStack2[T any] struct {
	head *node[T]
}

func NewInfStack2[T any]() *InfStack2[T] {
	return &InfStack2[T]{
		head: nil,
	}
}

func (s *InfStack2[T]) Push(item T) error {
	node := &node[T]{
		item: item,
		next: s.head,
	}
	s.head = node
	return nil
}

func (s *InfStack2[T]) Pop() (T, error) {
	if s.head == nil {
		var item T
		return item, ErrStackIsEmpty
	}

	item := s.head
	s.head = s.head.next
	return item.item, nil
}

func (s *InfStack2[T]) Iterator() gof.Iterator[T] {
	return &infStackIterator2[T]{
		stack:   s,
		current: s.head,
	}
}

type infStackIterator2[T any] struct {
	stack   *InfStack2[T]
	current *node[T]
}

func (i *infStackIterator2[T]) HasNext() bool {
	return i.current != nil
}

func (i *infStackIterator2[T]) Next() (T, error) {
	var item T
	if !i.HasNext() {
		return item, gof.ErrIteratorNext
	}
	item = i.current.item
	i.current = i.current.next
	return item, nil
}
