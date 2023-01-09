/*
Runtime
16 ms
Beats
76.81%

Memory
5.1 MB
Beats
89.49%
*/
package main

func main() {}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	min, max := -1, -1
	a1, a2 := len(nums1), len(nums2)
	l := a1 + a2
	num := 1.0
	if l%2 == 0 {
		min = l/2 - 1
		max = l / 2
		num = 2
	} else {
		max = l / 2
	}

	current := 0
	i, j := 0, 0
	sum, val := 0, 0
	for i < a1 || j < a2 {
		if i >= a1 {
			val = nums2[j]
			j++
		} else if j >= a2 {
			val = nums1[i]
			i++
		} else if nums1[i] >= nums2[j] {
			val = nums2[j]
			j++
		} else {
			val = nums1[i]
			i++
		}
		if current == min || current == max {
			sum += val
		}
		current++
	}
	return float64(sum) / num
}
