/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Balanced Binary Tree.
Memory Usage: 5.8 MB, less than 74.89% of Go online submissions for Balanced Binary Tree.
*/
package main

func main() {}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := check(root)
	return ok
}

// check checks the tree from a given parent and returns a height of the tree
// if the tree is balanced. Otherwise return false as the second return value.
func check(parent *TreeNode) (int, bool) {
	if parent == nil {
		return 0, true
	}

	l, ok := check(parent.Left)
	if !ok {
		return 0, false
	}
	r, ok := check(parent.Right)
	if !ok {
		return 0, false
	}

	if diff := abs(l - r); diff > 1 {
		return 0, false
	}
	return max(l, r) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
