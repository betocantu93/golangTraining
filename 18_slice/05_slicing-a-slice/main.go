package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 10)

	for i := 0; i < 10; i++ {

		mySlice = append(mySlice, i)

	}

	fmt.Print("[1:2] ")
	fmt.Println(mySlice[1:2])
	fmt.Print("[:2] ")
	fmt.Println(mySlice[:2])
	fmt.Print("[5:] ")
	fmt.Println(mySlice[5:])
	fmt.Print("[:] ")
	fmt.Println(mySlice[:])

}
