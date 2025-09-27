package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// stop immediately
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			jobCtx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			doWork(jobCtx, job)
			cancel()
		}
	}
}

func doWork(ctx context.Context, j Job) {
	select {
	case <-time.After(100 * time.Millisecond):
		// finished
		fmt.Println("Job finished for ID:", j.ID)
	case <-ctx.Done():
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workerCount := runtime.NumCPU()

	fmt.Println("worker count:", workerCount)

	jobs := make(chan Job, 64)
	var wg sync.WaitGroup

	// adding worker
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// producer
	go func() {
		defer close(jobs)
		for i := 0; i < 1000; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- Job{ID: i}:
			}
		}
	}()

	// shut down trigger
	time.AfterFunc(2*time.Second, func() {
		log.Println("cancel now")
		cancel()
	})

	wg.Wait()
}
