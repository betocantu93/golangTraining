package main

import "fmt"

func main() {
	word := "Hello"
	c := rune(word[0])
	letter := string(c)

	fmt.Printf("%T - %v \n", c, c)
	fmt.Printf("%T - %v \n", letter, letter)

}
