// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_containsNearbyDuplicate() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	// Output:
	// true
	// true
	// false
}
