package main

import "fmt"

func main() {
	fmt.Printf("Testing Add: %d\n", add(5, 6))
}

func add(a, b int) int {
	return a + b
}
