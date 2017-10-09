package main

import "fmt"

func main() {

	fmt.Println(hashFunction("Go", 12))

}

func hashFunction(word string, buckets int) int {

	letter := int(word[0])
	bucket := letter % buckets

	return bucket

}
