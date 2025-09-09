package main

import "fmt"

func main() {
	worker := 10
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < worker; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				c <- fmt.Sprint(id, ": ", j)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < worker; i++ {
			<-done
		}
		close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}
}
