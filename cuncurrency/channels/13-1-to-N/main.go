package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c := make(chan int)
	done := make(chan bool)

	worker := 10

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < worker; i++ {
		go func() {
			for val := range c {
				fmt.Println(val)
			}
			done <- true
		}()
	}

	for i := 0; i < worker; i++ {
		<-done
	}

	fmt.Println(time.Since(start))
}
