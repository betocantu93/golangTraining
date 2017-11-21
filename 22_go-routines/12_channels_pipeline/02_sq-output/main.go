package main

import "fmt"

func main() {

	for n := range sq(sq(sq(gen(5, 6, 7, 1, 23, 1, 3, 2, 1, 3, 5, 5, 10)))) {
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
