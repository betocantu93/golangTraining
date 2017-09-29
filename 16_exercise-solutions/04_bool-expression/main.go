package main

import "fmt"

func main() {

	fmt.Println(testExpression())

}

func testExpression() bool {

	return (true && false) || (false && true) || !(false && false)

}
