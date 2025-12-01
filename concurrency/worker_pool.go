package concurrency

import (
	"errors"
	"sync"
)

var ErrShutdown = errors.New("worker pool has been shut down")

type Task[R any] interface {
	Execute() (R, error)
}

type WorkerPool[T Task[R], R any] interface {
	Run() error
	Submit(task T) error
	Shutdown() error
}

type workerPool[T Task[R], R any] struct {
	jobs    chan T
	results chan R
	errors  chan error

	done chan struct{}

	errorHandler  func(error)
	resultHandler func(R)

	workerCount int
	wg          sync.WaitGroup
}

type NewWorkerPoolArgs[T Task[R], R any] struct {
	WorkerCount   int
	ErrorHandler  func(error)
	ResultHandler func(R)
}

func NewWorkerPool[T Task[R], R any](args NewWorkerPoolArgs[T, R]) WorkerPool[T, R] {
	return &workerPool[T, R]{
		jobs:          make(chan T),
		results:       make(chan R),
		errors:        make(chan error),
		done:          make(chan struct{}),
		errorHandler:  args.ErrorHandler,
		resultHandler: args.ResultHandler,
		workerCount:   args.WorkerCount,
	}
}

func (wp *workerPool[T, R]) Run() error {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for {
				select {
				case task := <-wp.jobs:
					result, err := task.Execute()
					if err != nil {
						wp.errorHandler(err)
						continue
					}
					wp.results <- result
				case <-wp.done:
					return
				}
			}
		}()
	}

	ch := make(chan struct{})
	go func() {
		for r := range wp.results {
			wp.resultHandler(r)
		}
		close(ch)
	}()

	wp.wg.Wait()

	close(wp.results)

	<-ch

	return ErrShutdown
}

func (wp *workerPool[T, R]) Submit(task T) error {
	wp.jobs <- task
	return nil
}

func (wp *workerPool[T, R]) Shutdown() error {
	close(wp.done)
	wp.wg.Wait()
	return nil
}
