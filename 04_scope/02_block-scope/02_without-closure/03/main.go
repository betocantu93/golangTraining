package main

import "fmt"

var h func() int

func main() {
	x := 0

	increment := func() int {
		x++
		fmt.Println(x)
		return x
	}

	increment()
	increment()
	increment()



	h = increment

	another()
}

func another() {
	h()
}
