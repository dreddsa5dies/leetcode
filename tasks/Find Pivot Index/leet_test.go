// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_pivotIndex() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
	// Output:
	// 3
	// -1
	// 0
}
