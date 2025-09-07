package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addS(a, b string) string {
	return a + b
}

type addConstraint interface {
	int | float64 | string
}

func addT[T addConstraint](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(4, 5))
	fmt.Println(addF(5.5, 5))
	fmt.Println(addS("Aj", "ay"))
	fmt.Println("-----------------------------")
	fmt.Println(addT(4, 5))
	fmt.Println(addT(5.5, 5))
	fmt.Println(addT("Aj", "ay"))
}
