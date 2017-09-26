package main

import "fmt"

func main() {

	apply(func(x interface{}) {
		fmt.Println(x)
	}, 1, "Hola", 4.3, "como estas", 99)

}

func apply(callback func(interface{}), args ...interface{}) {

	for _, i := range args {
		callback(i)
	}

}
