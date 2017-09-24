package main

import "fmt"

func main() {

	for i := 50; i <= 10000; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

}
