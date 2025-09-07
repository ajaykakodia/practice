package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

func hello(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "LaLa Agent",
		},
		ltk: true,
	}

	p2 := person{
		first: "Ajay",
	}

	// p2.speak()
	// sa1.speak()

	hello(p2)
	hello(sa1)
}
