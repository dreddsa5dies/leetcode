// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_isIsomorphic() {
	fmt.Println(isIsomorphic("1", "a"))
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	// Output:
	// true
	// true
	// false
	// true
}
