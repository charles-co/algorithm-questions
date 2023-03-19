package tests

import (
	"github.com/charles-co/algorithm_questions/trees"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tt := []struct {
		name      string
		postorder []int
		inorder   []int
		expected  *trees.TreeNode
	}{
		{
			name:      "Test 1",
			postorder: []int{9, 15, 7, 20, 3},
			inorder:   []int{9, 3, 15, 20, 7},
			expected: &trees.TreeNode{
				Val: 3,
				Left: &trees.TreeNode{
					Val: 9,
				},
				Right: &trees.TreeNode{
					Val: 20,
					Left: &trees.TreeNode{
						Val: 15,
					},
					Right: &trees.TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name:      "Test 2",
			postorder: []int{-1},
			inorder:   []int{-1},
			expected: &trees.TreeNode{
				Val: -1,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := trees.BuildTree(tc.postorder, tc.inorder)

			if !trees.IsSameTree(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
