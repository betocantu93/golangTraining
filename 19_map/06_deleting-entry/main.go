package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"Roberto": 10,
		"Rodrigo": 11,
		"Rúben":   17,
	}

	m["Filomeno"] = 30

	for w, v := range m {

		fmt.Println(w, " ", v)

	}

	m["filomeno"] = 10
	m["Filomeno"] = 15

	delete(m, "filomeno")

	fmt.Println(m)

}
