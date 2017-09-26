package main

import "fmt"

func main() {

	s, x := greet("Alberto ", "Cantu")

	fmt.Println(s, x)

}

func greet(name, father_name string) (string, string) {

	return fmt.Sprint(name, father_name), fmt.Sprint(name, father_name)

}
