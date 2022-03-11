package main

func main() {}

/*
Runtime: 72 ms, faster than 93.04% of Go online submissions for Container With Most Water.
Memory Usage: 8.7 MB, less than 88.71% of Go online submissions for Container With Most Water.
*/
func Max(x, y int) int {
	if x >= y {
		return x
	}

	return y
}

func maxArea(height []int) int {

	// length of input array
	size := len(height)

	// tow pointers, left init as 0, right init as size-1
	left, right := 0, size-1

	// maximal width between leftmost stick and rightmost stick
	maxWidth := size - 1

	// area also known as the amount of water
	area := 0

	// trade-off between width and height
	// scan each possible width and compute maximal area
	for width := maxWidth; width > 0; width-- {

		if height[left] < height[right] {

			// the height of lefthand side is shorter
			area = Max(area, width*height[left])

			// update left index to righthand side
			left += 1

		} else {

			// the height of righthand side is shorter
			area = Max(area, width*height[right])

			// update left index to righthand side
			right -= 1
		}

	}
	return area
}

// // прямой перебор - time out, оно и понятно, O(n^2)!
// // если упростить (вынести правый край) за цикл for, то будет O(n)
// func maxArea(height []int) int {
// 	ma := 0
// 	for i := 0; i <= len(height)-1; i++ {
// 		for k := len(height); k >= i; k-- {
// 			t := square(height[i:k])
// 			if t > ma {
// 				ma = t
// 			}
// 		}
// 	}
// 	return ma
// }

// // x - left, y - length, z - right
// func square(area []int) int {
// 	if len(area) == 0 {
// 		return 0
// 	}
// 	sq := 0
// 	x, y, z := area[0], len(area)-1, area[len(area)-1]
// 	if x >= z {
// 		sq = z * y
// 	} else {
// 		sq = x * y
// 	}
// 	return sq
// }
