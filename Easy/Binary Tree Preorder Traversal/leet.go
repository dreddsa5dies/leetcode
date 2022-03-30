package main

import "sort"

func main() {}

// definition for a binary tree node.
type treenode struct {
    val int
    left *treenode
    right *treenode
}
/*
Runtime: 3 ms, faster than 29.78% of Go online submissions for Binary Tree Preorder Traversal.
Memory Usage: 2 MB, less than 42.08% of Go online submissions for Binary Tree Preorder Traversal.
*/
func preorderTraversal(root *TreeNode) []int {
	fin := []int{}

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
			if node != nil {
					preorder(node.Left)

					fin = append(fin, node.Val)

					preorder(node.Right)
			}
	}

	preorder(root)

	return sort.Ints(fin)
}

