package main

import "sort"

type s []string

func main() {

	group := []string{"Alberto", "Juan", "Diego", "Rodrigo", "Gustavo", "Yehuda", "Bill"}
	sort.Sort(sort.StringSlice(group))

}
