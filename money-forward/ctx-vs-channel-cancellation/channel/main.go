package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs { //exits when channel is closed and drained
		doWork(j)
	}
}

func doWork(j Job) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Job finished :", j.ID)
}

func main() {
	jobs := make(chan Job, 64)

	workerCount := runtime.NumCPU()

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Multi-producer safe close pattern:
	var prodWg sync.WaitGroup
	startProducer := func(from, to int) {
		prodWg.Add(1)
		go func() {
			defer prodWg.Done()
			for i := from; i < to; i++ {
				jobs <- Job{ID: i}
			}
		}()
	}

	startProducer(0, 500)
	startProducer(500, 1000)

	// close jobs once all producers finish
	go func() {
		prodWg.Wait()
		close(jobs)
	}()

	wg.Wait()
}
