package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {

	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

func info(z Shape) {

	fmt.Printf("%T", z)
	fmt.Println(z)
	fmt.Println(z.area())

}

func main() {

	s := Square{
		side: 10,
	}

	c := Circle{
		radius: 5,
	}

	info(s)
	info(c)

}
