package main

func main() {}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 321 ms, faster than 26.91% of Go online submissions for Minimum Depth of Binary Tree.
Memory Usage: 21.4 MB, less than 55.92% of Go online submissions for Minimum Depth of Binary Tree.
*/
// recursion
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	return min(left, right) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
Runtime: 378 ms, faster than 9.28% of Go online submissions for Minimum Depth of Binary Tree.
Memory Usage: 21.1 MB, less than 61.95% of Go online submissions for Minimum Depth of Binary Tree.
*/
// BFS
// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	q := []*TreeNode{root}
// 	out := 1
// 	for len(q) != 0 {
// 		former := len(q)
// 		for i := 0; i < former; i++ {
// 			if q[i].Left == nil && q[i].Right == nil {
// 				return out
// 			}
// 			if q[i].Left != nil {
// 				q = append(q, q[i].Left)
// 			}
// 			if q[i].Right != nil {
// 				q = append(q, q[i].Right)
// 			}
// 		}
// 		q = q[former:]
// 		out++
// 	}
// 	return -1
// }
