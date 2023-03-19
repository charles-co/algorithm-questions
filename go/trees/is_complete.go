package trees

func IsComplete(root *TreeNode) bool {

	if root == nil {
		return false
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {

		var next []*TreeNode
		null := false

		for _, node := range queue {

			if node.Left != nil {
				if null {
					return false
				}

				if node.Right == nil {
					if len(next)&1 != 0 {
						return false
					}
				}
				next = append(next, node.Left)
			}

			if node.Right != nil {

				if node.Left == nil || null {
					return false
				}
				next = append(next, node.Right)
			}

			if node.Left == nil && node.Right == nil {
				null = true
			}
		}
		queue = next

	}
	return true
}
