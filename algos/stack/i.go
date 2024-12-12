package stack

import (
	"errors"
	"sandbox/gof"
)

var (
	ErrStackIsEmpty = errors.New("stack is empty")
)

type Stack[T any] interface {
	Push(item T) error
	Pop() (T, error)
	Iterator() gof.Iterator[T]
}
