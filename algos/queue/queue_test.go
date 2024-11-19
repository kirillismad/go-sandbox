package queue

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sandbox/gof"
	"testing"
	"time"
)

func TestQueueEnque(t *testing.T) {
	cases := []struct {
		name  string
		queue Queue[int]
	}{
		{
			name:  "SliceQueue",
			queue: NewSliceQueue[int](),
		},
		{
			name:  "ListQueue",
			queue: NewListQueue[int](),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("c.name: %v\n", c.name)
			err := c.queue.Enque(rand.Intn(100))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestQueueDeque(t *testing.T) {
	cases := []struct {
		name  string
		queue Queue[int]
	}{
		{
			name:  "SliceQueue",
			queue: NewSliceQueue[int](),
		},
		{
			name:  "ListQueue",
			queue: NewListQueue[int](),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("c.name: %v\n", c.name)
			err := c.queue.Enque(rand.Intn(100))
			if err != nil {
				t.Error(err)
			}
			v, err := c.queue.Deque()
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("v: %v\n", v)
		})
	}
}

func TestQueueIter(t *testing.T) {
	const count = 10
	cases := []struct {
		name  string
		queue Queue[int]
	}{
		{
			name:  "SliceQueue",
			queue: NewSliceQueue[int](),
		},
		{
			name:  "ListQueue",
			queue: NewListQueue[int](),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("c.name: %v\n", c.name)
			for i := 0; i < count; i++ {
				err := c.queue.Enque(i)
				if err != nil {
					t.Error(err)
				}
			}

			var iter gof.Iterator[int]
			switch q := c.queue.(type) {
			case *SliceQueue[int]:
				iter = q.Iterator()
			case *ListQueue[int]:
				iter = q.Iterator()
			}
			for {
				v, err := iter.Next()
				if err != nil {
					if errors.Is(err, gof.ErrIteratorNext) {
						break
					}
					t.Error(err)
				}
				fmt.Printf("v: %v\n", v)
			}
		})
	}
}

func TestSliceQueue(t *testing.T) {
	cases := []struct {
		name  string
		queue Queue[int]
	}{
		{
			name:  "SliceQueue",
			queue: NewSliceQueue[int](),
		},
		{
			name:  "ListQueue",
			queue: NewListQueue[int](),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("c.name: %v\n", c.name)
			const timeout = 2 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			go func() {
				for i := 0; true; i++ {
					select {
					case <-ctx.Done():
						return
					default:
						c.queue.Enque(i)
					}
					time.Sleep(time.Duration(rand.Intn(int(timeout.Milliseconds()/20))) * time.Millisecond)
				}
			}()
			go func() {
				n := int(timeout.Milliseconds() / 10)
				after := time.After(time.Duration(timeout.Nanoseconds() / 2))
				for {
					select {
					case <-ctx.Done():
						return
					case <-after:
						n /= 4
					default:
						v, _ := c.queue.Deque()
						fmt.Printf("v: %v\n", v)
					}
					time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
				}
			}()
			<-ctx.Done()
		})
	}
}
