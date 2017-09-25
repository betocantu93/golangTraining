package main

import "fmt"

func main() {

	var number1 float64
	var number2 float64

	fmt.Println("Ingresa un número")

	fmt.Scan(&number1)

	fmt.Println("Ingresa el otro")

	fmt.Scan(&number2)

	fmt.Println("el resultado es: ", number1/number2)

}
