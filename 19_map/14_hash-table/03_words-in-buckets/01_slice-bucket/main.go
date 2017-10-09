package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")

	if err != nil {
		panic("error")
	}

	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)

	for i := 0; i <= 12; i++ {
		buckets = append(buckets, []string{})
	}

	for scanner.Scan() {

		word := scanner.Text()
		n := HashBucket(word, 12)

		buckets[n] = append(buckets[n], word)

	}

	for i, _ := range buckets {

		fmt.Println(buckets[i])
	}

	//fmt.Println(buckets)

}

func HashBucket(word string, buckets int) int {

	letter := int(word[0])

	return letter % buckets

}
