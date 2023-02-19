package trees

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func ZigZagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		var level []int
		var next []*TreeNode
		for _, node := range queue {
			level = append(level, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, level)
		queue = next
	}

	for i := 0; i < len(res); i++ {
		if i%2 == 1 {
			reverse(res[i])
		}
	}

	return res
}
