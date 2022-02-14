/*
Runtime: 37 ms, faster than 63.54% of Go online submissions for Reverse String.
Memory Usage: 7.1 MB, less than 25.37% of Go online submissions for Reverse String.
*/
package main

func main() {}

func reverseString(s []byte) {
	var b []byte
	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}
	copy(s, b)
}
