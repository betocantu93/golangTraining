package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		panic("error")
	}

	bs, _ := ioutil.ReadAll(res.Body)

	str := string(bs)

	fmt.Println(str)
}
