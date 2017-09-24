package main

import "fmt"

func main() {

	if !true {
		fmt.Println("This will not run")
	}

	if !false {
		fmt.Println("This will run")
	}
}
