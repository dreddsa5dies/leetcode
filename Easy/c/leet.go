/*
Runtime: 5 ms, faster than 48.43% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.3 MB, less than 86.50% of Go online submissions for Maximum Depth of Binary Tree.
*/
package main

func main() {}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lD := maxDepth(root.Left)
	rD := maxDepth(root.Right)

	if lD > rD {
		return lD + 1
	} else {
		return rD + 1
	}
}
