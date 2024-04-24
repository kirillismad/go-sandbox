package stack

import (
	"errors"
	"sandbox/gof"
)

var (
	ErrStackisFull = errors.New("stack is full")
)

type FixedStack[T any] struct {
	current int
	slice   []T
}

func NewFixedStack[T any](size int) *FixedStack[T] {
	s := new(FixedStack[T])
	s.current = 0
	s.slice = make([]T, size)
	return s
}

func (s *FixedStack[T]) Push(item T) error {
	if s.current < len(s.slice) {
		s.slice[s.current] = item
		s.current++
		return nil
	}
	return ErrStackisFull
}

func (s *FixedStack[T]) Pop() (T, error) {
	if s.current == 0 {
		var item T
		return item, ErrStackIsEmpty
	}

	s.current--
	return s.slice[s.current], nil
}

func (s *FixedStack[T]) Iterator() gof.Iterator[T] {
	return &FixedStackIterator[T]{
		stack: s,
		index: s.current,
	}
}

type FixedStackIterator[T any] struct {
	stack *FixedStack[T]
	index int
}

func (i *FixedStackIterator[T]) HasNext() bool {
	return i.index > 0
}

func (i *FixedStackIterator[T]) Next() (T, error) {
	if !i.HasNext() {
		var item T
		return item, gof.ErrIteratorNext
	}
	i.index--
	item := i.stack.slice[i.index]
	return item, nil
}
