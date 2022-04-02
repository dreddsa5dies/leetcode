/*
Runtime: 32 ms, faster than 60.80% of Go online submissions for First Unique Character in a String.
Memory Usage: 5.5 MB, less than 61.77% of Go online submissions for First Unique Character in a String.
*/
package main

import (
	"sort"
	"strings"
)

func main() {}

// если 1 символ он и есть уникальный, если нет, то пилим в мапу все руны со счетчиком, находим индексы уникальных рун, сортируем и возвращаем наименьший
func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	char := map[rune]int{}

	for _, v := range s {
		char[v]++
	}

	indexes := []int{}

	for k, v := range char {
		if v == 1 {
			indexes = append(indexes, strings.IndexRune(s, k))
		}
	}

	if len(indexes) > 0 {
		sort.Ints(indexes)
		return indexes[0]
	} else {
		return -1
	}
}

// остальные решения работают с ASCII, не с UTF-8, типа:
// Runtime: 7 ms, faster than 89.47% of Go online submissions for First Unique Character in a String.
// Memory Usage: 5.3 MB, less than 82.41% of Go online submissions for First Unique Character in a String.
// массив работает быстрее, нет аллокаций в мапу, не надо сортировать.
/*
func firstUniqChar(s string) int {
    m := make([]int, 26)
    for _, v := range s {
        m[int(v- 'a')]++
    }
    for i, v := range s {
        if m[int(v- 'a')] == 1 {
            return i
        }
    }
    return -1
}
*/
