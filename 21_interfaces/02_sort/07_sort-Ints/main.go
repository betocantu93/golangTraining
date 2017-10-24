package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{5, 4, 3, 1, 1, 3}

	sort.Ints(n)
	fmt.Println(n)

}
