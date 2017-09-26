package main

import "fmt"

func main() {

	var n int64

	fmt.Println("Please enter a number for factorial.")

	fmt.Scan(&n)

	fmt.Println(factorial(n))

}

func factorial(n int64) int64 {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}
