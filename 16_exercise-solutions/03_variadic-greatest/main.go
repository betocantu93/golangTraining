package main

import "fmt"

func main() {

	biggest := greatest(3, 4, 2, 1, 5, 199, 23, 1, 23, 25, 3, 2, 1)

	fmt.Println(biggest)

}

func greatest(l ...int) int {

	var greater int

	for _, n := range l {
		if n > greater {
			greater = n
		}
	}
	return greater
}
