/*
Runtime: 11 ms, faster than 47.04% of Go online submissions for Roman to Integer.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Roman to Integer.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("1211"))
}

func romanToInt(s string) int {
	ro := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	count := 0
	for i := 0; i < len(s); i++ {
		switch {
		case len(s) == 0:
			return count
		case len(s) == 1:
			return ro[string(s[0])]
		case string(s[i]) == "I" && i+1 <= len(s)-1 && string(s[i+1]) == "V":
			count += 4
			i++
		case string(s[i]) == "I" && i+1 <= len(s)-1 && string(s[i+1]) == "X":
			count += 9
			i++
		case string(s[i]) == "X" && i+1 <= len(s)-1 && string(s[i+1]) == "L":
			count += 40
			i++
		case string(s[i]) == "X" && i+1 <= len(s)-1 && string(s[i+1]) == "C":
			count += 90
			i++
		case string(s[i]) == "C" && i+1 <= len(s)-1 && string(s[i+1]) == "D":
			count += 400
			i++
		case string(s[i]) == "C" && i+1 <= len(s)-1 && string(s[i+1]) == "M":
			count += 900
			i++
		default:
			count += ro[string(s[i])]
		}
	}
	return count
}
