package main

import "fmt"

func main() {

	var name string
	var apellido string
	var indirect *string = &apellido

	fmt.Println("Ingresa un nombre")

	fmt.Scanf("%s %s", &name, &apellido)


	switch name {

	case "Rosita":
		fmt.Println("Wassup Rossita")
	case "Alberto":
		fmt.Println("Wassup Alberto")
	case "Rodrigo":
		fmt.Println("Wassup Rodrigo")
	default:
		fmt.Println("Wassup stranger")
	}

	fmt.Printf("%v %v \n",indirect, *indirect)
	*indirect = "Salazar"
	fmt.Printf("%v %v \n",indirect, apellido)

}
