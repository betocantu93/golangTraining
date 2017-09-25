package main

import "fmt"

func main() {

	var sum int

	var limit int

	fmt.Println("Enter a limit")

	fmt.Scan(&limit)

	for i := 0; i < limit; i++ {

		if i%3 == 0 || i%5 == 0 {
			sum += i
		}

	}

	fmt.Println(sum)
}
