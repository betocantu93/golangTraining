package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world!")
}

func mexico() {
	fmt.Println("Mexico")
}

func main() {
	defer world()
	defer mexico()
	hello()
}
