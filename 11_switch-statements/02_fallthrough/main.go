package main

import "fmt"

func main() {

	var name string

	fmt.Scanf("%s", &name)

	switch name {

	case "Rosita":
		fmt.Println("Wassup Rossita")
	case "Alberto":
		fmt.Println("Wassup Alberto")
	case "Rodrigo":
		fmt.Println("Wassup Rodrigo")
		fallthrough
	default:
		fmt.Println("Wassup stranger")
	}

}
