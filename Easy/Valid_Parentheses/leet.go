/*
Runtime: 7 ms, faster than 7.01% of Go online submissions for Valid Parentheses.
Memory Usage: 7.5 MB, less than 6.33% of Go online submissions for Valid Parentheses.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(true, isValid(""))
	fmt.Println(false, isValid("(]"))
	fmt.Println(true, isValid("()"))
	fmt.Println(true, isValid("()[]{}"))
	fmt.Println(false, isValid("([)]"))
}

func isValid(s string) bool {
	x := true

	valids := []string{"()", "{}", "[]"}

	count := 0

	for _, val := range valids {
		if strings.Contains(s, val) {
			s = strings.Replace(s, val, "", -1)
			count++
		}
	}

	if s == "" {
		return x
	} else if count == 0 {
		x = false
	} else {
		return isValid(s)
	}

	return x
}
