package concurrent

import (
	"context"
	"sync"
)

// Task is a function that returns an error.
type Task func() error

// Parallel runs the given tasks in parallel with a maximum number of workers.
// It returns the first error encountered, or nil if all tasks completed successfully.
// If stopOnError is true, it returns the first error encountered and stops processing.
//
// The implementation uses a worker pool to limit the number of concurrent tasks.
// It supports context cancellation and deadlines to stop processing early.
func Parallel(ctx context.Context, tasks []Task, maxWorkers int, stopOnError ...bool) error {
	tasksCh := make(chan Task, len(tasks))
	resultsCh := make(chan error, len(tasks))
	stopOnErr := false
	if len(stopOnError) > 0 {
		stopOnErr = stopOnError[0]
	}

	// Use a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(maxWorkers)

	// Start worker goroutines.
	for i := 0; i < maxWorkers; i++ {
		go func() {
			defer wg.Done()
			worker(tasksCh, resultsCh)
		}()
	}

	// Send tasks to workers.
	for _, task := range tasks {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case tasksCh <- task:
		}
	}
	close(tasksCh)

	// Wait for all tasks to complete.
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	var taskErr error

loop:
	// Read results.
	for {
		select {
		case <-ctx.Done():
			return ctx.Err() // Context canceled or deadline exceeded.
		case err, ok := <-resultsCh:
			// Drain the results channel, but return the first error.
			if !ok {
				break loop // resultsCh has been closed.
			}
			if stopOnErr && err != nil {
				return err
			}

			// Save the first error.
			if taskErr == nil && err != nil {
				taskErr = err
			}

		}
	}

	// All tasks completed successfully.
	return taskErr
}

// worker runs tasks from tasks channel and sends the result to results channel.
func worker(tasks <-chan Task, results chan<- error) {
	for task := range tasks {
		results <- task()
	}
}
