package main

import "fmt"

func main() {

	myGreeting := map[int]string{

		0: "Hola",
		1: "Hello",
		2: "Bonjour",
	}

	fmt.Println(myGreeting)

	//delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {

		fmt.Println("That value exists")
		fmt.Println("val", val)
		fmt.Println("exists", exists)

	} else {

		fmt.Println("That value doesn't exists")
		fmt.Println("val", val)
		fmt.Println("exists", exists)

	}

	fmt.Println(myGreeting)

}
