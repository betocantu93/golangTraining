package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"Yehuda", "Katz"}
	sort.Strings(s)
	fmt.Println(s)
}
