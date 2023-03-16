package trees

func BuildTree(postorder []int, inorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(postorder) == 1 {
		return root
	}

	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = BuildTree(postorder[:i], inorder[:i])
	root.Right = BuildTree(postorder[i:len(postorder)-1], inorder[i+1:])

	return root

}

func BuildTreeRecursive(postorder []int, inorder []int) *TreeNode {
	
	rootIdx := len(postorder) - 1
	inorderMap := make(map[int]int)

	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func(lo, high int) *TreeNode

	build = func(lo, high int) *TreeNode {
		if lo > high {
			return nil
		}

		root := &TreeNode{Val: postorder[rootIdx]}
		rootIdx--

		root.Right = build(inorderMap[root.Val] + 1, high)
		root.Left = build(lo, inorderMap[root.Val] - 1)

		return root
	}


	return build(0, len(inorder)-1)	

}