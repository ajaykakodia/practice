package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1, c2 := make(chan int), make(chan int)
	go populate(c1)

	//go fanOutIn(c1, c2)

	go fanOutWithLimitedWorker(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			defer wg.Done()
			c2 <- timeConsumingWork(v2)
		}(v)
	}
	wg.Wait()
	close(c2)
}

func fanOutWithLimitedWorker(c1, c2 chan int) {
	var wg sync.WaitGroup

	const worker = 10
	wg.Add(10)
	for i := 0; i < worker; i++ {
		go func(id int) {
			defer wg.Done()
			for v := range c1 {
				fmt.Println("Id:", id, "val:", v)
				c2 <- timeConsumingWork(v)
			}
		}(i)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
