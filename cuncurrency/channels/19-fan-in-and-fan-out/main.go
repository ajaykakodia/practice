package main

import (
	"fmt"
)

func main() {
	c := gen(2, 3, 4, 5, 6, 7, 8, 9)

	c1 := sq(c)
	c2 := sq(c)
	c3 := sq(c)

	for sqs := range merge(c1, c2, c3) {
		fmt.Println(sqs)
	}

	ch := gen(10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	worker := 3
	cs := make([]<-chan int, worker)
	for i := 0; i < worker; i++ {
		cs[i] = sq(ch)
	}

	fmt.Println("---------------------------")

	for sqs := range merge(cs...) {
		fmt.Println(sqs)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range ch {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	// var wg sync.WaitGroup
	done := make(chan bool)
	chCount := len(chs)

	for _, ch := range chs {
		// wg.Add(1)
		go func(c <-chan int) {
			// defer wg.Done()
			for sqs := range c {
				out <- sqs
			}
			done <- true
		}(ch)
	}

	go func() {
		// wg.Wait()
		for i := 0; i < chCount; i++ {
			<-done
		}
		close(out)
	}()

	return out
}
