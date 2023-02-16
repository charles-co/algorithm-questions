package tests

import (
	"testing"

	"github.com/charles-co/algorithm_questions/dfs"
)

func TestMaxDepth(t *testing.T) {
	tt := []struct {
		name     string
		root     *dfs.TreeNode
		expected float64
	}{
		{name: "Test 1", root: &dfs.TreeNode{Val: 3, Left: &dfs.TreeNode{Val: 9}, Right: &dfs.TreeNode{Val: 20, Left: &dfs.TreeNode{Val: 15}, Right: &dfs.TreeNode{Val: 7}}}, expected: 3},
		{name: "Test 2", root: &dfs.TreeNode{Val: 1, Left: &dfs.TreeNode{Val: 2}, Right: &dfs.TreeNode{Val: 3, Left: &dfs.TreeNode{Val: 4}, Right: &dfs.TreeNode{Val: 5}}}, expected: 3},
		{name: "Test 3", root: &dfs.TreeNode{Val: 1, Left: &dfs.TreeNode{Val: 2, Left: &dfs.TreeNode{Val: 3, Left: &dfs.TreeNode{Val: 4, Left: &dfs.TreeNode{Val: 5}}}}}, expected: 5},
		{name: "Test 4", root: &dfs.TreeNode{Val: 1, Right: &dfs.TreeNode{Val: 2, Right: &dfs.TreeNode{Val: 3, Right: &dfs.TreeNode{Val: 4, Right: &dfs.TreeNode{Val: 5}}}}}, expected: 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := dfs.MaxDepth(tc.root)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}

}
