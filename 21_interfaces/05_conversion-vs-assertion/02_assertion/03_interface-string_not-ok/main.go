package main

import "fmt"

func main() {
	var name interface{} = 12

	str, ok := name.(string)
	if ok {
		fmt.Println(str)
		fmt.Printf("%T", name)
	} else {
		fmt.Println("Error")
	}
}
