package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		panic("error")
	}

	page, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		panic("Hola")
	}

	fmt.Printf("%s", page)

}
