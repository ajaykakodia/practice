package main

import "fmt"

func main() {
	e, o, q := make(chan int), make(chan int), make(chan int)

	go send(e, o, q)

	receive(e, o, q)

	fmt.Println("about to exit.....")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case val, ok := <-e:
			if ok {
				fmt.Println("Even:", val)
			}
		case val, ok := <-o:
			if ok {
				fmt.Println("Odd", val)
			}
		case val := <-q:
			fmt.Println(val)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
	close(q)
}
