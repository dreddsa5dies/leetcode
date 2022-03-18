// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_divide() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	// Output:
	// 3
	// -2
}
