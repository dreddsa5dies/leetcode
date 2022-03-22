// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_addBinary() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	// Output:
	// 100
	// 10101
}
