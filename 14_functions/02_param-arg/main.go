package main

import "fmt"

func main() {
	greet("Carlos")
	greet("Hugo")
}

func greet(name string) {

	fmt.Println("Hello", name)

}
