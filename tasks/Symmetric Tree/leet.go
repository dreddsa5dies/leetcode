/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
Memory Usage: 2.9 MB, less than 91.69% of Go online submissions for Symmetric Tree.
*/
package main

func main() {}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if isEmpty(root) {
		return true
	}

	return mirror(root.Left, root.Right)
}

func isEmpty(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return false
}

func mirror(root1, root2 *TreeNode) bool {
	if isEmpty(root1) && isEmpty(root2) {
		return true
	} else if isEmpty(root1) || isEmpty(root2) {
		return false
	}

	return root1.Val == root2.Val && mirror(root1.Left, root2.Right) && mirror(root1.Right, root2.Left)
}
