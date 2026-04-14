package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	ID       int
	Duration time.Duration
}

type Result struct {
	JobID    int
	TimedOut bool
}

type Aggregator struct {
	startedAt time.Time

	totalTasks  int
	successIDs  []int
	timeoutIDs  []int
	elapsedTime time.Duration
}

func (a *Aggregator) Aggregate(ch <-chan Result) {
	for result := range ch {
		a.totalTasks++
		if result.TimedOut {
			a.timeoutIDs = append(a.timeoutIDs, result.JobID)
			continue
		}
		a.successIDs = append(a.successIDs, result.JobID)
	}
	a.elapsedTime = time.Since(a.startedAt)
}

func (a Aggregator) PrintReport() {
	fmt.Printf("Total tasks: %d\n", a.totalTasks)
	fmt.Printf("Successful: %d\n", len(a.successIDs))
	fmt.Printf("Timeouts: %d\n", len(a.timeoutIDs))
	fmt.Printf("Successful IDs: %v\n", a.successIDs)
	fmt.Printf("Timeout IDs: %v\n", a.timeoutIDs)
	fmt.Printf("Total elapsed: %s\n", a.elapsedTime.Round(time.Millisecond))
}

func parseArgs(args []string) (int, []time.Duration, error) {
	if len(args) < 2 {
		return 0, nil, fmt.Errorf("usage: go run ./practice <workers> <task_ms_1> <task_ms_2> ...")
	}

	workers, err := strconv.Atoi(args[0])
	if err != nil || workers <= 0 {
		return 0, nil, fmt.Errorf("workers must be a positive integer")
	}

	durations := make([]time.Duration, 0, len(args)-1)
	for i, raw := range args[1:] {
		d, convErr := strconv.Atoi(raw)
		if convErr != nil || d < 0 {
			return 0, nil, fmt.Errorf("invalid duration at position %d: %q (must be non-negative integer in ms)", i+1, raw)
		}
		durations = append(durations, time.Duration(d)*time.Millisecond)
	}

	return workers, durations, nil
}

func main() {
	workers, durations, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	jobs := make(chan Job)
	results := make(chan Result)

	go func() {
		defer close(jobs)
		for i, d := range durations {
			jobs <- Job{ID: i + 1, Duration: d}
		}
	}()
	startedAt := time.Now()

	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for job := range jobs {
				select {
				case <-time.After(job.Duration):
					results <- Result{JobID: job.ID, TimedOut: false}
				case <-time.After(800 * time.Millisecond):
					results <- Result{JobID: job.ID, TimedOut: true}
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	a := &Aggregator{startedAt: startedAt}
	a.Aggregate(results)
	a.PrintReport()
}
