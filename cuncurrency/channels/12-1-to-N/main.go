package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for val := range c {
			fmt.Println(val)
		}
		done <- true
	}()

	go func() {
		for val := range c {
			fmt.Println(val)
		}
		done <- true
	}()

	<-done
	<-done
	fmt.Println(time.Since(start))
}
