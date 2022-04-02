/*
Runtime: 128 ms, faster than 38.62% of Go online submissions for Contains Duplicate.
Memory Usage: 8.4 MB, less than 94.12% of Go online submissions for Contains Duplicate.
*/
package main

import "sort"

func main() {}

func containsDuplicate(nums []int) bool {
	cd := false
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			cd = true
			break
		}
	}

	return cd
}

/*
// сортируя в функции срез, я изменяю его исходное состояние, что не очень хорошо, но зато в лучшем случае серьезно ускоряет поиск дупликатов, можно еще так:
func containsDuplicate(nums []int) bool {
	s := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := s[nums[i]]; ok {
			return true
		}
		// причем struct будет меньше bool по размеру выделяемой памяти
		s[nums[i]] = struct{}{}
	}
	return false
}
*/
