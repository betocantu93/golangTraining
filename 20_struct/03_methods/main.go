package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"Alberto", "Cantu", 40}
	p2 := person{"Rodrigo", "Altamirano", 8}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
