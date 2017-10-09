package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func (p DoubleZero) salute() (s string) {

	if p.LicenseToKill {
		s = "I can kill you now"
	} else {
		s = "I can't kill you"
	}

	return
}

func main() {

	p1 := DoubleZero{
		Person: Person{
			First: "Alberto",
			Last:  "Cantu",
			Age:   48,
		},
		First:         "Yahuda",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Carmen",
			Last:  "Hinojosa",
			Age:   12,
		},
		First:         "Tom",
		LicenseToKill: false,
	}

	p3 := DoubleZero{
		Person{
			"Miramar",
			"Santillar",
			12,
		},
		"Stephen",
		false,
	}

	fmt.Printf("%T - %v and %v \n", p1, p1, p1.salute())
	fmt.Printf("%T - %v  and %v \n", p2, p2, p2.salute())
	fmt.Printf("%T - %v  and %v \n", p3, p3, p3.salute())

}
