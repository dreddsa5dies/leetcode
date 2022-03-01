/*
Runtime: 12 ms, faster than 97.89% of Go online submissions for Power of Three.
Memory Usage: 6.2 MB, less than 100.00% of Go online submissions for Power of Three.
*/
package main

func main() {}

// тут определяем граничные значения и проверяем на делимость на 3 до конца
// и это быстрее, чем через строки
func isPowerOfThree(n int) bool {
	if n%3 == 0 && n > 0 {
		for n%3 == 0 {
			n /= 3
		}
	}
	return n == 1
}

// без прохождения по циклу - через преобразование в строку
/*
Runtime: 16 ms, faster than 93.16% of Go online submissions for Power of Three.
Memory Usage: 6.2 MB, less than 90.53% of Go online submissions for Power of Three.
*/
/*
func isPowerOfThree(n int) bool {
	s := strconv.FormatInt(int64(n), 3)
	return s[0] == '1' && len(s) == strings.Count(s, "0")+1
}
*/
