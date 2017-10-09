package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (p Person) fullName() (fullName string) {

	fullName = p.First + " " + p.Last

	return

}

func (dz DoubleZero) Greeting() {
	fmt.Println(dz.Person.fullName(), "so good to see you")
}

func main() {
	p1 := Person{
		"Bill",
		"Gates",
		66,
	}

	p2 := DoubleZero{
		Person{
			"Yehuda",
			"Katz",
			50,
		},
		false,
	}

	p4 := DoubleZero{
		p1,
		false,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
	p4.Greeting()

}
