package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{4, 3, 1, 23, 31, 12, 3, 1, 23, 3, 5}

	sort.Sort(sort.IntSlice(n))

	fmt.Println(n)

}
