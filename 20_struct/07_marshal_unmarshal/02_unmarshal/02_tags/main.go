package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {

	var p1 Person

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"Yehuda", "Last": "Katz", "wisdom score": 45}`)

	json.Unmarshal(bs, &p1)

	fmt.Println("__________________")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)

}
