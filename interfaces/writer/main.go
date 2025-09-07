package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	p := person{
		first: "LaLa Here",
	}

	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatalf("failed to create error %s", err)
	}
	defer f.Close()

	s := []byte("Hello Gophers! ")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("Failed to write into the file %s", err)
	}

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers!!")
	fmt.Println(b.String())
	b.Reset()

	b.WriteString("Hello I am here ")
	b.Write([]byte("Gophers!!"))
	fmt.Println(b.String())

	var nb bytes.Buffer

	p.writeOut(f)
	p.writeOut(&nb)
	fmt.Println(nb.String())
}
