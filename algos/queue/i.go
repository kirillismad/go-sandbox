package queue

import (
	"errors"
)

var (
	ErrQueueIsEmpty = errors.New("queue is empty")
)

type Queue[T any] interface {
	Enque(item T) error
	Deque() (T, error)
}
