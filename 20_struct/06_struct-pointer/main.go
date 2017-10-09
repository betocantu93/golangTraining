package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {

	p1 := &person{"Rodrigo Medina", 90}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)

}
