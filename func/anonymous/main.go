package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("Anonymous function called with my name:", s)
	}("Lala Here")

	y := bar()
	r := y()
	fmt.Println(r)

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

func foo() int {
	fmt.Println("Foo ran")
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
