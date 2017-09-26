package main

import "fmt"

func main() {

	s := greet("Alberto ", "Cantu")

	fmt.Println(s)

}

func greet(name, father_name string) string {

	return fmt.Sprint(name, father_name)

}
