package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Printf("%s \n", "Hello World")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
