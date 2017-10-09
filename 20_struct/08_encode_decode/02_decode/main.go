package main

import (
	"encoding/json"
	"strings"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person

	rdr := strings.NewReader(`{"First":"Bill", "Last":"Gates", "Age":65}`)

	dc := json.NewDecoder(rdr)

	err := dc.Decode(&p1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
