package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main(){


	 book := "A most significant digit (MSD) radix sort can be used to sort keys in lexicographic order. Unlike a least significant digit (LSD) radix sort, " +
		 "a most significant digit radix sort does not necessarily preserve the original order of duplicate keys. An MSD radix sort starts processing the keys" +
		 " from the most significant digit, leftmost digit, to the least significant digit, rightmost digit. This sequence is opposite that of least significant digit (LSD) " +
		 "radix sorts. An MSD radix sort stops rearranging the position of a key when the processing reaches a unique prefix of the key. Some MSD radix sorts use one " +
		 "level of buckets in which to group the keys. See the counting sort and pigeonhole sort articles. Other MSD radix sorts use multiple levels of buckets, " +
		 "which form a trie or a path in a trie. A postman's sort / postal sort is a kind of MSD radix sort."

	 scanner := bufio.NewScanner(strings.NewReader(book))

	 scanner.Split(bufio.ScanWords)

	 for scanner.Scan() {

	 	fmt.Println(scanner.Text())

	 }

	 if err := scanner.Err(); err != nil {
	 	fmt.Fprintln(os.Stderr, "reading input: ", err)
	 }

}