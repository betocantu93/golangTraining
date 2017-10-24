package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Linus", "Bill", "Steve", "Elon", "Wozniak", "Mark", "Yehuda", "Tom"}
	sort.Sort(sort.Reverse(sort.StringSlice(s)))

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
