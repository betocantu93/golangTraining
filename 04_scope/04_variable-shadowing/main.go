package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {

	square := square(5)

	fmt.Println(square)

}
