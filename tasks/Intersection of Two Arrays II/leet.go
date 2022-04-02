package main

func main() {}

/*
Runtime: 8 ms, faster than 25.04% of Go online submissions for Intersection of Two Arrays II.
Memory Usage: 2.8 MB, less than 85.32% of Go online submissions for Intersection of Two Arrays II.
*/
// тут О(n^2) по сложности, но О(1) по памяти
func intersect(nums1 []int, nums2 []int) []int {
	index := 0
	len1 := len(nums1)
	len2 := len(nums2)

	for i := 0; i < len1; i++ {
		for j := index; j < len2; j++ {
			if nums1[i] == nums2[j] {
				nums1[index] = nums1[i]
				nums2[index], nums2[j] = nums2[j], nums2[index]
				index++
				break
			}
		}
	}

	return nums1[:index]
}

/*
// тут уже с созданием среза
func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
loop:
	for _, n1 := range nums1 {
        for j, n2 := range nums2 {
			if n1 == n2 {
				res = append(res, n2)
				if len(nums2) > 0 {
					nums2 = append(nums2[:j], nums2[j+1:]...)
				} else {
					break loop
				}
				break
			}
		}
	}
	return res
}
*/
