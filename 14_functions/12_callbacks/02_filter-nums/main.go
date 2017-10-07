package main

import "fmt"

func filterInts(callback func(int) bool, numbers []int) []int {

	xs := []int{}

	for _, n := range numbers {

		if callback(n) {
			xs = append(xs, n)
		}

	}

	return xs
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	xs := filterInts(func(n int) bool {
		return n > 5
	}, arr)

	fmt.Println(xs)

	x := filter(func(x interface{}) bool {
		if num, ok := x.(int); ok {
			return num > 15
		} else {
			return x == "Hola"
		}
	}, 1, "Hola", 4.3, "como estás", 99)

	fmt.Println(x)

}

func filter(callback func(interface{}) bool, arr ...interface{}) []interface{} {

	var xs []interface{}

	for _, v := range arr {

		if callback(v) {
			xs = append(xs, v)
		}

	}

	return xs
}
