/*
Runtime: 8 ms, faster than 40.72% of Go online submissions for Fizz Buzz.
Memory Usage: 4.1 MB, less than 30.24% of Go online submissions for Fizz Buzz.
*/
package main

import "strconv"

func main() {}

func fizzBuzz(n int) []string {
	fin := []string{}
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fin = append(fin, "FizzBuzz")
		case i%3 == 0:
			fin = append(fin, "Fizz")
		case i%5 == 0:
			fin = append(fin, "Buzz")
		default:
			fin = append(fin, strconv.Itoa(i))
		}
	}
	return fin
}
