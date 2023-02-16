package dfs

// Definition for a binary tree node.

func IsSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return check(p.Left, q.Right) && check(p.Right, q.Left)
}
