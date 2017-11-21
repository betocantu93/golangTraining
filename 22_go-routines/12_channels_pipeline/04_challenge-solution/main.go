package main

import "fmt"

func main() {

	in := gen()

	c := factorial(in)

	for n := range c {
		fmt.Println(n)
	}

}

func gen() <-chan int {
	out := make(chan int)

	go func() {

		for i := 3; i <= 103; i++ {
			out <- i % 13
		}

		close(out)
	}()

	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		for n := range c {

			out <- fact(n)
		}

		close(out)

	}()

	return out
}

func fact(n int) int {
	total := 1

	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}
