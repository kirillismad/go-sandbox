package gof

import "errors"

var (
	ErrIteratorNext = errors.New("there is no next item")
)

type Iterator[T any] interface {
	Next() (T, error)
	HasNext() bool
}
