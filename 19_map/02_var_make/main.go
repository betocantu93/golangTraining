package main

import "fmt"

func main() {

	m := map[string]string{"car1": "red", "car2": "blue"}

	for k, v := range m {

		fmt.Println(k, " ", v)

	}

	f := map[float64]int{4.78: 3, 7.89: 2}

	for k, v := range f {

		fmt.Println(k, " ", v)

	}

	var g map[string]int

	fmt.Println(g == nil)

}
