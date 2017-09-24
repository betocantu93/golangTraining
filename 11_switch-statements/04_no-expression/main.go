package main

import "fmt"

func main() {

	var name string

	fmt.Scanf("%s", &name)

	switch {
	case name[0] == 'M' || name[0] == 'm':
		fmt.Println("Your name starts with M!")
	default:
		fmt.Println("This is the default")
	}

}
