package main

import "fmt"

func main() {

	n := 10
	c := factorial(n)

	fmt.Println("The factorial of", n, "is", <-c)

}

func factorial(n int) chan int {

	out := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {

			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}
