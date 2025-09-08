package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	var counter int64
	const gc = 100
	var wg sync.WaitGroup

	wg.Add(gc)
	for i := 0; i < gc; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("Final Counter: \t", counter)
}
