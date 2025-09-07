package main

import "fmt"

func main() {
	nums := []int{3, 5, 7, 5, 2}
	fmt.Println(foo(nums...))
}

func foo(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}
