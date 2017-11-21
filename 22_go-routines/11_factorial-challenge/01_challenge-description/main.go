package main

import (
	"fmt"
)

func main() {

	n := 10

	for i := 0; i <= 100; i++ {
		f := factorial(n)
		fmt.Println("The factorial of", n, "is", f)
	}

}

func factorial(n int) int {
	acum := 1
	for i := 1; i <= n; i++ {
		acum *= i
	}

	return acum
}
