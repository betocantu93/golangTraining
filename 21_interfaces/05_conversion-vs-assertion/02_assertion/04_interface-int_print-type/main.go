package main

import "fmt"

func main() {
	var name interface{} = 12
	str, ok := name.(int)
	if ok {
		fmt.Println(str)
		fmt.Printf("%T\n", name)
	} else {
		fmt.Println("Error")
	}
}
