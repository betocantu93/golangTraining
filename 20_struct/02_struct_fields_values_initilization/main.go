package main

import "fmt"

type person struct {
	first string `json:first_name`
	last string
	age int
}

func main() {
	var p3 person
	p1 := person{"Alberto", "Cantu", 45}
	p2 := person{"Rodrigo", "Altamirano", 67}

	persons := make(map[int]person)

	persons[0] = p1
	persons[1] = p2

	fmt.Println(persons)
	fmt.Printf("%T - %v - %v - %v \n", p3, p3.first, p3.last, p3.age)


}