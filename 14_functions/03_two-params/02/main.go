package main

import "fmt"

func main() {
	greet("Carlos", "Sanchez")
	greet("Hugo", "Martinez")
}

func greet(name, father_name string) {

	fmt.Println("Hello", name, father_name)

}
