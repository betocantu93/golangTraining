package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name     string `json:"nombre"`
	LastName string
	Age      int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p1 := Person{
			Name:     "Alberto",
			LastName: "Cantu",
			Age:      24,
		}
		json.NewEncoder(w).Encode(p1)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
