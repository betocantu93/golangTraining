package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		panic("error")
	}

	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 12)

	for scanner.Scan() {

		n := HashBucket(scanner.Text(), len(buckets))
		buckets[n]++

	}

	fmt.Println(buckets)

}

func HashBucket(word string, buckets int) (bucket int) {

	letter := int(word[0])
	bucket = letter % buckets

	return

}
