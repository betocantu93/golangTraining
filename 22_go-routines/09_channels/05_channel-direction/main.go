package main

import "fmt"

func main() {

	c := incrementor()
	cSum := puller(c)

	for n := range cSum {
		//fmt.Printf("%T %T\n",n, cSum)
		fmt.Println(n)
	}
}

func incrementor() <-chan int {

	out := make(chan int)

	go func() {
		for i := 0; i <= 100; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()

	return out

}
