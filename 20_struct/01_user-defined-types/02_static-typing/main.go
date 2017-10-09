package main

import "fmt"

type foo int

func main() {

	var myVar foo


	myVar = 42

	fmt.Printf("%T %v \n", myVar, myVar)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	//fmt.Println(myVar + yourAge)

	fmt.Println(foo(yourAge) + myVar)
	fmt.Println(int(myVar) + yourAge)

}