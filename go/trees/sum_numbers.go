package trees

func SumNumbers(root *TreeNode) int {
	return sumNumbers(root, 0)
}

func sumNumbers(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return sumNumbers(root.Left, sum) + sumNumbers(root.Right, sum)
}