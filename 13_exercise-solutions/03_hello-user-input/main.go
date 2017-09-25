package main

import "fmt"

func main() {

	var name string
	var apellido string
	var apellido2 int64

	fmt.Println("Ingresa tu nombre")

	fmt.Scanf("%v %v %v", &name, &apellido, &apellido2)

	fmt.Println("Hello", name, " ", apellido, " ", apellido2)
}
