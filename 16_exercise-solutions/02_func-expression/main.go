package main

import "fmt"

func main() {

	f := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	b, flag := f(10)

	fmt.Println(b, " ", flag)
}

func divideAndEven(n int) (int, bool) {

	return n / 2, n%2 == 0
}
