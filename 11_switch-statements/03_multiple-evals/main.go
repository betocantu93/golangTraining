package main

import "fmt"

func main() {

	var name string

	fmt.Scanf("%s", &name)

	switch name {

	case "Rosita", "Juan":
		fmt.Println("Wassup Rossita, err, Juan?")
	case "Alberto":
		fmt.Println("Wassup Alberto")
	case "Rodrigo":
		fmt.Println("Wassup Rodrigo")
		fallthrough
	default:
		fmt.Println("Wassup stranger")
	}

}
