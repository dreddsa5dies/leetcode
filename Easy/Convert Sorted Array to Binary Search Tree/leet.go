/*
Runtime: 7 ms, faster than 20.43% of Go online submissions for Convert Sorted Array to Binary Search Tree.
Memory Usage: 3.5 MB, less than 42.74% of Go online submissions for Convert Sorted Array to Binary Search Tree.
*/

package main

func main() {}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	return sortedArrayToBSTUtil(nums, 0, length-1)
}

func sortedArrayToBSTUtil(nums []int, first int, last int) *TreeNode {
	if first > last {
		return nil
	}

	if first == last {
		return &TreeNode{
			Val: nums[first],
		}
	}

	mid := (first + last) / 2

	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = sortedArrayToBSTUtil(nums, first, mid-1)
	root.Right = sortedArrayToBSTUtil(nums, mid+1, last)
	return root
}
