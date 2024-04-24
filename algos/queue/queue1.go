package queue

import (
	"sandbox/gof"
	"sync"
)

type Queue1[T any] struct {
	slice []T
	head  int
	tail  int
	size  int
	mutex *sync.Cond
}

const initLen = 4

func NewQueue1[T any]() *Queue1[T] {
	return &Queue1[T]{
		slice: make([]T, initLen),
		head:  0,
		tail:  0,
		size:  0,
		mutex: sync.NewCond(&sync.Mutex{}),
	}
}

func (q *Queue1[T]) resize(size int) {
	newSlice := make([]T, size)

	j := q.head
	for i := 0; i < q.size; i++ {
		newSlice[i] = q.slice[j]

		j++
		if j == len(q.slice) {
			j -= len(q.slice)
		}
	}
	q.slice = newSlice
	q.head = 0
	q.tail = q.size
}

func (q *Queue1[T]) Enque(item T) error {
	q.mutex.L.Lock()
	defer q.mutex.L.Unlock()

	if q.size == len(q.slice) {
		q.resize(len(q.slice) * 2)
	}
	q.slice[q.tail] = item

	q.tail++
	if q.tail == len(q.slice) {
		q.tail -= len(q.slice)
	}

	q.size++
	q.mutex.Signal()
	return nil
}

func (q *Queue1[T]) Deque() (T, error) {
	q.mutex.L.Lock()
	defer q.mutex.L.Unlock()

	for q.size == 0 {
		q.mutex.Wait()
	}

	item := q.slice[q.head]
	q.head++
	if q.head == len(q.slice) {
		q.head -= len(q.slice)
	}
	q.size--

	if len(q.slice) >= initLen*2 && len(q.slice)/q.size == 4 {
		q.resize(len(q.slice) / 2)
	}
	return item, nil
}

func (q *Queue1[T]) Iterator() gof.Iterator[T] {
	return &queue1Iterator[T]{
		queue: q,
		index: 0,
	}
}

type queue1Iterator[T any] struct {
	queue *Queue1[T]
	index int
}

func (i *queue1Iterator[T]) HasNext() bool {
	return i.index < i.queue.size
}

func (i *queue1Iterator[T]) Next() (T, error) {
	var item T
	if !i.HasNext() {
		return item, gof.ErrIteratorNext
	}
	index := i.queue.head + i.index
	if index >= len(i.queue.slice) {
		index = index - len(i.queue.slice)
	}
	item = i.queue.slice[index]
	i.index++
	return item, nil
}
