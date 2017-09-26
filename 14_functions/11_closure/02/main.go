package main

import "fmt"

var x int = 0

func main() {

	fmt.Println(increment())
	fmt.Println(increment())

}

func increment() int {

	x++
	return x

}
