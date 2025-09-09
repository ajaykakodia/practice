package main

import "fmt"

func main() {
	c := incrementor()
	pc := puller(c)
	for val := range pc {
		fmt.Println(val)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for val := range c {
			sum += val
		}
		out <- sum
		close(out)
	}()
	return out
}
