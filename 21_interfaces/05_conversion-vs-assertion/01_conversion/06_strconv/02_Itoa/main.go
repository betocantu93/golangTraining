package main

import (
	"strconv"
	"fmt"
)

func main(){
	var x = 12
	var y = "I have this many: " + strconv.Itoa(x)

	fmt.Println(y)

}