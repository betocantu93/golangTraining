package main

import "fmt"

type foo int

func main() {

	var myVar foo

	myVar = 42

	fmt.Printf("%T %v \n", myVar, myVar)

}