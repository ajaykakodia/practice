package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subs)
	fmt.Printf("%T\n", doWhatIsAskedFor)

	x := doWhatIsAskedFor(654, 234, add)
	fmt.Println(x)

	y := doWhatIsAskedFor(324, 123, subs)
	fmt.Println(y)
}

func add(a, b int) int {
	return a + b
}

func subs(a, b int) int {
	return a - b
}

func doWhatIsAskedFor(a, b int, f func(int, int) int) int {
	return f(a, b)
}
