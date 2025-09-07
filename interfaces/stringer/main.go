package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("this is the book: ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("this number is ", strconv.Itoa(int(c)))
}

func logInfo(f fmt.Stringer) {
	log.Println("Additional Log Info Here:", f.String())
}

func main() {
	b := book{
		title: "Good to Know",
	}

	var c count = 42
	c++

	//fmt.Println(b)
	//fmt.Println(c)

	//log.Println(b)
	//log.Println(c)

	logInfo(b)
	logInfo(c)
}
