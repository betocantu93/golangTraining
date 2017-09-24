package main

import "fmt"

func main() {

	for i := 50; i <= 10000; i++ {
		fmt.Println(i, " - ", string(i), []byte(string(i)))
	}

}
