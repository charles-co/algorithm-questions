package tests

import (
	"testing"

	"github.com/charles-co/algorithm_questions/trees"
)

func TestMaxDepth(t *testing.T) {
	tt := []struct {
		name     string
		root     *trees.TreeNode
		expected float64
	}{
		{name: "Test 1", root: &trees.TreeNode{Val: 3, Left: &trees.TreeNode{Val: 9}, Right: &trees.TreeNode{Val: 20, Left: &trees.TreeNode{Val: 15}, Right: &trees.TreeNode{Val: 7}}}, expected: 3},
		{name: "Test 2", root: &trees.TreeNode{Val: 1, Left: &trees.TreeNode{Val: 2}, Right: &trees.TreeNode{Val: 3, Left: &trees.TreeNode{Val: 4}, Right: &trees.TreeNode{Val: 5}}}, expected: 3},
		{name: "Test 3", root: &trees.TreeNode{Val: 1, Left: &trees.TreeNode{Val: 2, Left: &trees.TreeNode{Val: 3, Left: &trees.TreeNode{Val: 4, Left: &trees.TreeNode{Val: 5}}}}}, expected: 5},
		{name: "Test 4", root: &trees.TreeNode{Val: 1, Right: &trees.TreeNode{Val: 2, Right: &trees.TreeNode{Val: 3, Right: &trees.TreeNode{Val: 4, Right: &trees.TreeNode{Val: 5}}}}}, expected: 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := trees.MaxDepth(tc.root)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}

}
