package main

import "fmt"

func zero(x int) {
	x = 10
}

func main() {

	x := 5

	fmt.Println(x)

	zero(x)

	

	fmt.Println(x)

}
