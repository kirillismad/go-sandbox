package concurrency

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type taskAdd struct {
	a, b int
}

func (t taskAdd) Execute() (int, error) {
	time.Sleep(200 * time.Millisecond)
	return t.a + t.b, nil
}

func Test_WorkerPool(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		wp := NewWorkerPool(NewWorkerPoolArgs[taskAdd, int]{
			WorkerCount: 2,
			ErrorHandler: func(err error) {
				t.Errorf("unexpected error: %v", err)
			},
			ResultHandler: func(result int) {
				t.Logf("result: %d", result)
			},
		})

		go func() {
			if err := wp.Run(); err != nil && !errors.Is(err, ErrShutdown) {
				t.Errorf("failed to start worker pool: %v", err)
			}
		}()

		const taskCount = 10
		tasks := make([]taskAdd, 0, taskCount)
		for i := range taskCount {
			tasks = append(tasks, taskAdd{a: i, b: i * 2})
		}
		for _, task := range tasks {
			if err := wp.Submit(task); err != nil {
				t.Errorf("failed to submit task: %v", err)
			}
		}

		require.NoError(t, wp.Shutdown())
	})
}
