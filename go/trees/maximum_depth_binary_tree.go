package trees

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
import (
	"math"
)

func MaxDepth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return 1 + math.Max(float64(MaxDepth(root.Left)), float64(MaxDepth(root.Right)))
}
