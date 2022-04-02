/*
Runtime: 41 ms, faster than 36.81% of Go online submissions for Move Zeroes.
Memory Usage: 7.1 MB, less than 57.76% of Go online submissions for Move Zeroes.
*/
package main

func main() {}

// создаем срез по размеру, отбираем объекты не равные 0, копируем срез в срез, т.к. нет возврата
func moveZeroes(nums []int) {
	tmp := make([]int, len(nums))
	j := 0
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] != 0 {
			tmp[j] = nums[i]
			j++
		}
	}
	copy(nums, tmp)
}

/*
// тут без дополнительной памяти
func moveZeroes(nums []int)  {
    l := 0

    for r, _ := range nums {
        if nums[r] != 0 {
            nums[l], nums[r] = nums[r], nums[l]
            l++
        }
    }
}
*/
