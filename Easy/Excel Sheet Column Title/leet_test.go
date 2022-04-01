// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_convertToTitle() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	// Output:
	// A
	// AB
	// ZY
}
