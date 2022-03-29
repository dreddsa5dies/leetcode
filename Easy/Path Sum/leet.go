package main

func main() {}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 6 ms, faster than 53.92% of Go online submissions for Path Sum.
Memory Usage: 4.7 MB, less than 76.96% of Go online submissions for Path Sum.
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
