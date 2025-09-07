package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking.")
}

func (d *dog) run() {
	fmt.Println("My name is", d.first, "and I am running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{
		first: "Tommy",
	}

	d1.walk()
	d1.run()
	// cannot use d1 (variable of struct type dog) as youngin value in argument to youngRun:
	// dog does not implement youngin (method run has pointer receiver)
	// youngRun(d1)

	d2 := &dog{
		first: "Kalu",
	}
	d2.walk()
	d2.run()
	youngRun(d2)
}
