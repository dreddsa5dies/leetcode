package main

import "math"

func main() {}

/*
Runtime: 3 ms, faster than 57.36% of Go online submissions for Reverse Integer.
Memory Usage: 2.1 MB, less than 51.91% of Go online submissions for Reverse Integer.
*/
// правильный подход
func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return result
}

/*
Runtime: 3 ms, faster than 57.36% of Go online submissions for Reverse Integer.
Memory Usage: 2.3 MB, less than 28.00% of Go online submissions for Reverse Integer.
*/
// // прямой подход, зато свой: О(n^2)
// func reverse(x int) int {
// 	var retv int
// 	if x >= 0 {
// 		retv = rev(x)
// 	} else {
// 		retv = -rev(-x)
// 	}

// 	if retv > 2147483647 || retv < -2147483648 {
// 		return 0
// 	} else {
// 		return retv
// 	}
// }

// func rev(x int) int {
// 	tmp := []int{}
// 	counter := x
// 	a := 10
// 	for i := 1; i <= counter; i *= 10 {
// 		tmp = append(tmp, x%a)
// 		x = (x - x%a) / 10
// 	}

// 	fin := 0
// 	a = int(math.Pow10(len(tmp) - 1))
// 	for i := 0; i <= len(tmp)-1; i++ {
// 		fin += tmp[i] * a
// 		a = a / 10
// 	}

// 	return fin
// }
