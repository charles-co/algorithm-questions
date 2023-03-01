package tests

import (
	"github.com/charles-co/algorithm_questions/trees"
	"testing"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	tt := []struct {
		name     string
		input    *trees.TreeNode
		expected []*trees.TreeNode
	}{
		{
			name: "Example 1",
			input: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val: 2,
					Left: &trees.TreeNode{
						Val: 4,
					},
				},
				Right: &trees.TreeNode{
					Val: 3,
					Left: &trees.TreeNode{
						Val: 2,
						Left: &trees.TreeNode{
							Val: 4,
						},
					},
					Right: &trees.TreeNode{
						Val: 4,
					},
				},
			},
			expected: []*trees.TreeNode{
				{
					Val: 2,
					Left: &trees.TreeNode{
						Val: 4,
					},
				},
				{
					Val: 4,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := trees.FindDuplicateSubtrees(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %d, got %d", len(tc.expected), len(actual))
			}
			for i := range actual {
				if actual[i].Val != tc.expected[i].Val {
					t.Errorf("Expected %d, got %d", tc.expected[i].Val, actual[i].Val)
				}
			}
		})
	}
}
