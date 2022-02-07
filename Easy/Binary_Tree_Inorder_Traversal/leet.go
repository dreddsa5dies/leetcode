/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2 MB, less than 87.04% of Go online submissions for Binary Tree Inorder Traversal.
*/
package main

import "fmt"

func main() {
	fmt.Println()
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)

			result = append(result, node.Val)

			inorder(node.Right)
		}
	}

	inorder(root)

	return result
}
