package stack

import (
	"sandbox/gof"
	"sync"
)

type InfStack[T any] struct {
	slice   []T
	current int
	mutex   *sync.Mutex
}

const initLen = 4

func NewInfStack[T any]() *InfStack[T] {
	return &InfStack[T]{
		slice:   make([]T, initLen),
		current: 0,
		mutex:   new(sync.Mutex),
	}
}

func (s *InfStack[T]) Push(item T) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.current == len(s.slice) {
		s.resize(len(s.slice) * 2)
	}

	s.slice[s.current] = item
	s.current++
	return nil
}

func (s *InfStack[T]) Pop() (T, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.current == 0 {
		var item T
		return item, ErrStackIsEmpty
	}
	s.current--
	item := s.slice[s.current]
	defer func() {
		// stack is as twice bigger as initial one and filled only in a quarter
		if len(s.slice) >= initLen*2 && len(s.slice)/s.current == 4 {
			s.resize(len(s.slice) / 2)
		}
	}()
	return item, nil
}

func (s *InfStack[T]) resize(size int) {
	newSlice := make([]T, size)
	copy(newSlice, s.slice)
	s.slice = newSlice
}

func (s *InfStack[T]) Iterator() gof.Iterator[T] {
	return &infStackIterator[T]{
		stack: s,
		index: s.current,
	}
}

type infStackIterator[T any] struct {
	stack *InfStack[T]
	index int
}

func (i *infStackIterator[T]) Next() (T, error) {
	var item T
	if !i.HasNext() {
		return item, gof.ErrIteratorNext
	}
	i.index--
	item = i.stack.slice[i.index]
	return item, nil
}

func (i *infStackIterator[T]) HasNext() bool {
	return i.index > 0
}
