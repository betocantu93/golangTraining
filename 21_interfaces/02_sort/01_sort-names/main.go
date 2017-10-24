package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int { return len(p) }
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type MyNew interface {
	sum() int
}
type Ed struct {
	x int
	y int
}

func (e Ed) sum() int {
	return e.x + e.y
}
func hello(m MyNew) {
	fmt.Println(m.sum())
}

type hj []string

func main() {

	m := Ed{x: 10, y: 20}
	hello(m)

	studyGroup := people{"Alberto", "José", "Ramiro", "Diego"}

	s := []string{"Alberto", "José", "Ramiro", "Diego"}

	i := []int{1, 23, 4, 5, 27, 3, 0}

	sort.Sort(studyGroup)
	sort.Sort(sort.StringSlice(s))
	sort.Sort(sort.IntSlice(i))

	//BEFORE REVERSE
	fmt.Println("BEFORE REVERSE")
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(i, "\n")

	//AFTER REVERSE
	fmt.Println("AFTER REVERSE")

	sort.Sort(sort.Reverse(studyGroup))
	sort.Sort(sort.Reverse(sort.IntSlice(i)))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))

	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(i)

}
