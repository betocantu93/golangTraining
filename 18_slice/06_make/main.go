package main

import "fmt"

func main() {

	myTransactions := make([][]int, 0, 5)

	for i := 0; i < 10; i++ {

		transaction := make([]int, 0)

		for j := i; j < 10; j++ {
			transaction = append(transaction, j+i)
		}

		myTransactions = append(myTransactions, transaction)

	}

	fmt.Printf("%T\n", myTransactions)
	fmt.Println(myTransactions)

}




