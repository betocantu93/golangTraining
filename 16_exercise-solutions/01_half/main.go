package main

import "fmt"

func main() {

	b, flag := divideAndEven(10)

	fmt.Println(b, " ", flag)
}

func divideAndEven(n int) (int, bool) {

	return n / 2, n%2 == 0
}
