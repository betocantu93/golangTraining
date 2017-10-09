package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		panic("error")
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {

		words[sc.Text()] = ""

	}

	for k, _ := range words {

		fmt.Println(k)

	}

}
