package trees

import (
	"strconv"
	"fmt"
)
func FindDuplicateSubtrees(root *TreeNode) (res []*TreeNode) {
	nodes := make(map[string][]*TreeNode)
	res = make([]*TreeNode, 0)

	var trv func(*TreeNode) string

	trv = func(node *TreeNode) string {
		if node == nil {
			return "nil"
		}
		key := strconv.Itoa(node.Val) + "," + trv(node.Left) + "," + trv(node.Right)
		nodes[key] = append(nodes[key], node)

		return key
	}
	trv(root)
	fmt.Println(nodes)

	for _, v := range nodes {
		if len(v) > 1 {
			res = append(res, v[0])
			fmt.Println(res)
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }

	return res
	
}