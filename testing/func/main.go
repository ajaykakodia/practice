package main

import "fmt"

func main() {
	fmt.Printf("Testing Add: %d\n", add(5, 6))
	fmt.Printf("2 + 3 = %d\n", mySum(2, 3))
	fmt.Printf("2 + 3 + 5 = %d\n", mySum(2, 3, 5))
	fmt.Printf("7 + 8 = %d\n", mySum(7, 8))
}

func add(a, b int) int {
	return a + b
}

func mySum(ns ...int) int {
	sum := 0
	for _, n := range ns {
		sum += n
	}

	return sum
}

func mySumWrong(ns ...int) int {
	sum := 0
	for _, n := range ns {
		sum += n
	}

	return sum + 1
}
