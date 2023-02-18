package dfs

// Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.

import (
	"math"
)

func MinDiffBST(root *TreeNode) int {
	min := math.MaxInt64
	var prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			if node.Val-prev.Val < min {
				min = node.Val - prev.Val
			}
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return min
}