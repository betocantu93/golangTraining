package main

import "fmt"

func main() {

	c := gen(5, 6, 7, 1, 23, 1, 3, 2, 1, 3, 5, 5, 10)

	for n := range sq(c) {
		fmt.Println(n)
	}

}

func gen(nums ...int) <-chan int {

	out := make(chan int)

	go func() {

		for _, n := range nums {
			out <- n
		}
		close(out)

	}()

	return out
}

func sq(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		for n := range in {
			out <- n * n
		}
		close(out)

	}()

	return out

}
