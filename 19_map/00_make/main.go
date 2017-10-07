package main

import "fmt"

func main() {

	myMap := make(map[string]string)

	myMap["jugador1"] = "Gaitan"
	myMap["jugador2"] = "Cristiano Ronaldo"

	fmt.Println("My Map: ", myMap)

	fmt.Println(myMap["jugador1"])

	v := myMap["jugador1"]

	delete(myMap, "jugador1")

	fmt.Println(v)

}
