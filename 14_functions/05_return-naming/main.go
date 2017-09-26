package main

import "fmt"

func main() {

	s := greet("Alberto ", "Cantu")

	fmt.Println(s)

}

func greet(name, father_name string) (s string) {

	s = fmt.Sprint(name, father_name)

	return

}
