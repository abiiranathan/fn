package concurrent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/abiiranathan/fn"
)

func TestParallel(t *testing.T) {
	type args struct {
		ctx         context.Context
		tasks       []Task
		maxWorkers  int
		stopOnError []bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				ctx:        context.Background(),
				tasks:      []Task{func() error { return nil }},
				maxWorkers: 1,
			},
			wantErr: false,
		},
		{
			name: "Test 2",
			args: args{
				ctx:        context.Background(),
				tasks:      []Task{func() error { return nil }},
				maxWorkers: 1,
			},
			wantErr: false,
		},
		{
			name: "Test 3",
			args: args{
				ctx:         context.Background(),
				tasks:       []Task{func() error { return fmt.Errorf("request timeout") }},
				maxWorkers:  1,
				stopOnError: []bool{true},
			},
			wantErr: true,
		},
		{
			name: "Test 4",
			args: args{
				ctx: context.Background(),
				tasks: []Task{
					func() error { return nil },
					func() error { fmt.Println("task 4 - 2"); return nil },
					func() error { fmt.Println("Task 4 - 3"); return nil },
				},
				maxWorkers:  3,
				stopOnError: []bool{true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Parallel(tt.args.ctx, tt.args.tasks, tt.args.maxWorkers, tt.args.stopOnError...); (err != nil) != tt.wantErr {
				t.Errorf("Parallel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContextCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	tasks := []Task{
		func() error {
			return nil
		},
		func() error {
			time.Sleep(time.Second * 4)
			return nil
		},
	}

	err := Parallel(ctx, tasks, 2)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Expected context.DeadlineExceeded error, got %v", err)
	}
}

func TestParallelSum(t *testing.T) {
	var sum int
	var lock sync.Mutex
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	chunks := fn.Chunk(arr, 3)
	tasks := make([]Task, 0, len(chunks))
	for _, chunk := range chunks {
		chunk := chunk
		tasks = append(tasks, func() error {
			var s int
			for _, v := range chunk {
				s += v
			}

			lock.Lock()
			sum += s
			lock.Unlock()
			return nil
		})
	}

	err := Parallel(context.Background(), tasks, 3)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if sum != 55 {
		t.Errorf("Expected sum to be 55, got %d", sum)
	}

}
