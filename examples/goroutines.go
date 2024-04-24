package examples

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func DemoGoroutins() {
	// runFuncsSync()
	// runFuncsAsync()
	// demoProducerWorker()
	// demoProcessAndGather()
	// demoServiceABC()
	// demoConcurrentMap1()
	demoConcurrentIncrement()
}

func runWithDelay(d time.Duration) {
	time.Sleep(d)
	fmt.Println("Print after:", d.Seconds(), "seconds")
}
func runFuncsSync() {
	fmt.Println(newline + fname(runFuncsSync))
	for i := 0; i < 10; i++ {
		// 1000 <= n <= 3000
		delay := rand.Intn(3000-1000+1) + 1000
		runWithDelay(time.Duration(delay) * time.Millisecond)
	}
}

func runFuncsAsync() {
	fmt.Println(newline + fname(runFuncsAsync))
	for i := 0; i < 10; i++ {
		// 1000 <= n <= 3000
		delay := rand.Intn(3000-1000+1) + 1000
		go runWithDelay(time.Duration(delay) * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
}

func producer(count int, ch chan<- int) {
	for i := 0; i < count; i++ {
		// 500 <= n <= 1000
		time.Sleep(time.Duration(rand.Intn(1000-500+1)+500) * time.Millisecond)
		ch <- i
		fmt.Println("P:", i)
	}
}

func worker(count int, ch <-chan int) {
	for i := 0; i < count; i++ {
		// 1000 <= n <= 2000
		time.Sleep(time.Duration(rand.Intn(2000-1000+1)+1000) * time.Millisecond)
		fmt.Println("W:", <-ch)
	}
}

func demoProducerWorker() {
	fmt.Println(newline + fname(demoProducerWorker))
	const count = 10
	ch := make(chan int, 3)

	go producer(count, ch)
	worker(count, ch)
}

func doubler(v int) int {
	time.Sleep(500 * time.Millisecond)
	return v * 2
}
func demoProcessAndGather() {
	const goro = 10
	ch := make(chan int, goro)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Println("IN:", i)
		}
		close(ch)
	}()
	result := processAndGather(ch, doubler, goro)
	fmt.Println(len(result))
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		i := i
		go func() {
			defer wg.Done()
			for v := range in {
				r := processor(v)
				out <- r
				fmt.Println("Worker:", i, "result:", r)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

type result1 int
type result2 int
type result3 int

func processService1(x int) (result1, error) {
	time.Sleep(450 * time.Millisecond)
	r := result1(x - 1)
	fmt.Println("S1:", r)
	// return r, nil\
	return result1(0), errors.New("s1 error")
}

func processService2(x int) (result2, error) {
	time.Sleep(450 * time.Millisecond)
	r := result2(x + 1)
	fmt.Println("S2:", r)
	return r, nil
}

func processService3(x1 result1, x2 result2) (result3, error) {
	time.Sleep(500 * time.Millisecond)
	r := result3(int(x1) * int(x2))
	fmt.Println("S3:", r)
	return r, nil
}

func demoServiceABC() {
	start := time.Now()

	input := 10

	out1 := make(chan result1, 1)
	out2 := make(chan result2, 1)
	r := make(chan result3, 1)
	errCh := make(chan error, 2)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		x1, err := processService1(input)
		if err != nil {
			errCh <- err
		} else {
			out1 <- x1
		}
		fmt.Println("G1")
	}()
	go func() {
		x2, err := processService2(input)
		if err != nil {
			errCh <- err
		} else {
			out2 <- x2
		}
		fmt.Println("G2")
	}()

	go func() {
		var x1 result1
		var x2 result2
		for i := 0; i < 2; i++ {
			select {
			case <-ctx.Done():
				// errCh <- ctx.Err()
				fmt.Println("G3")
				return
			case x1 = <-out1:
				continue
			case x2 = <-out2:
				continue
			}
		}
		result, err := processService3(x1, x2)
		if err != nil {
			errCh <- err
		} else {
			r <- result
		}
		fmt.Println("G3")
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Timeouted!", ctx)
	case result := <-r:
		fmt.Println("Result:", result)
	case err := <-errCh:
		fmt.Println("Error:", err)
	}

	fmt.Println("TimeConsumed:", time.Since(start))
	time.Sleep(3 * time.Second)
}
func demoConcurrentIncrement() {
	const workers = 10000
	var counter atomic.Int64
	counter.Store(42)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
			// atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter.Load())
}

func demoConcurrentMap1() {
	const workers = 10000
	m := make(map[int]int)
	var l sync.Mutex
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			l.Lock()
			defer l.Unlock()

			m[rand.Intn(10)]++
		}()
	}
	wg.Wait()
	fmt.Println(m)
}
