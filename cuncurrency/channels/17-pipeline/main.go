package main

import "fmt"

func main() {
	ch := gen(3, 4)
	sqc := sq(ch)
	fmt.Println(<-sqc) // 9
	fmt.Println(<-sqc) // 16

	for val := range sq(sq(gen(2, 3))) {
		fmt.Println(val)
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

func sq(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range c {
			out <- val * val
		}
		close(out)
	}()
	return out
}
