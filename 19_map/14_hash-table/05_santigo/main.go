package main

import "fmt"

func main() {

	var acum int
	var number int
	for i := 0; i <= 5; i++ {
		acum = 0
		for j := 0; j < 4; j++ {
			fmt.Println("Ingresa un número para el set: ", i+1)
			fmt.Scan(&number)
			acum += number
		}
		fmt.Println("Promedio del set: ", acum/4)
	}
}
