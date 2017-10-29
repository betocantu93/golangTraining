package main

import "fmt"

func main() {
	var name interface{} = 12
	fmt.Println(name.(int) + 10)
}
