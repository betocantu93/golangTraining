package main

import "fmt"

func main(){
	var name interface {} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Println(str)
	}else{
		fmt.Println("Error")
	}
}