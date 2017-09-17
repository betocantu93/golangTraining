package main

import "fmt"

var x = 5

func increment() int {
	x++
	return x

}

func main() {

	fmt.Println(increment())
	fmt.Println(increment())

}
