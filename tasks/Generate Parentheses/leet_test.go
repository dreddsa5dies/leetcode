// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_generateParenthesis() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
	// Output:
	// [((())) (()()) (())() ()(()) ()()()]
	// [()]
}
