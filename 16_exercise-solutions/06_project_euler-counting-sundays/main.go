package main

import "fmt"
import "math"

func main() {
	//fmt.Println()
	foo(5)
}

func foo(top float64) int {

	var resp int

	sum := int(math.Pow((top*(top+1))/2, 2))

	var acum int

	for i := 1; i <= int(top); i++ {
		acum += i * i
	}

	fmt.Println(sum - acum)

	return resp
}
