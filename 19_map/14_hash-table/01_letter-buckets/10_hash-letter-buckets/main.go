package main

import (
	"net/http"
	"bufio"
	"fmt"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		panic("error")
	}

	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 200)

	for scanner.Scan() {

		n := HashBucket(scanner.Text())
		buckets[n]++

	}

	fmt.Println(buckets[65:122])

}

func HashBucket(word string) int {

	return int(word[0])

}