/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
Memory Usage: 2 MB, less than 83.58% of Go online submissions for Pascal's Triangle.
*/
package main

func main() {}

func generate(numRows int) [][]int {
	fin := [][]int{}

	for line := 0; line < numRows; line++ {
		tmp := []int{}
		for i := 0; i <= line; i++ {
			tmp = append(tmp, binomialCoeff(line, i))
		}
		fin = append(fin, tmp)
	}
	return fin
}

func binomialCoeff(n, k int) int {
	res := 1

	if k > (n - k) {
		k = n - k
	}

	for i := 0; i < k; i++ {
		res *= n - i
		res /= i + 1
	}
	return res
}
