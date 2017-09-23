package main

import "fmt"

func main() {
	//x := 13 % 3

	for i := 1; i < 100; i++ {

		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}

	}
}
