package tests

import (
	"github.com/charles-co/algorithm_questions/trees"
	"testing"
)


func TestSumNumbers(t *testing.T){
	tt := []struct{
		name string
		root *trees.TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &trees.TreeNode{Val: 1, Left: &trees.TreeNode{Val: 2}, Right: &trees.TreeNode{Val: 3}},
			expected: 25,
		},
		{
			name: "Example 2",
			root: &trees.TreeNode{Val: 4, Left: &trees.TreeNode{Val: 9, Left: &trees.TreeNode{Val: 5}, Right: &trees.TreeNode{Val: 1}}, Right: &trees.TreeNode{Val: 0}},
			expected: 1026,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			actual := trees.SumNumbers(tc.root)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}