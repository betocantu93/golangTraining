package main

import "fmt"

func main() {
	c := factorial(4)

	for n := range c {
		fmt.Println(n)
	}

}

func factorial(num int) <-chan int {
	out := make(chan int)

	go func() {

		total := 1

		for i := num; i > 0; i-- {
			total *= i
		}

		out <- total
		close(out)

	}()

	return out
}
