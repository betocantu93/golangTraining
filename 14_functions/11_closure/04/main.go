package main

import "fmt"

func main() {

	incrementOne := wrapper()
	incrementTwo := wrapper()

	fmt.Println(incrementOne())
	fmt.Println(incrementOne())

	fmt.Println(incrementTwo())

}

func wrapper() func() int {

	var x int

	return func() int {
		x++
		return x
	}

}
