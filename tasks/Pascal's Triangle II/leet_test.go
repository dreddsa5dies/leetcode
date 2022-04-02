// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_getRow() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	// Output:
	// [1 3 3 1]
	// [1]
	// [1 1]
}
