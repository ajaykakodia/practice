// main package for exercises
package main

import (
	"fmt"
	"math"
)

type square struct {
	Side float64
}

func (s square) Area() float64 {
	return s.Side * s.Side
}

type rectangle struct {
	Length float64
	Width  float64
}

func (s rectangle) Area() float64 {
	return s.Length * s.Width
}

type circle struct {
	Radius float64
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type shape interface {
	Area() float64
}

func shapeArea(s shape) {
	a := s.Area()
	fmt.Printf("Area: %.2f\n", a)
}

func main() {
	sq := square{
		Side: 5,
	}

	rec := rectangle{
		Length: 4,
		Width:  5,
	}

	ci := circle{
		Radius: 3,
	}
	shapeArea(sq)
	shapeArea(rec)
	shapeArea(ci)
}
