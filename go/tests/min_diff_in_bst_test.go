package tests

import (
	"testing"

	"github.com/charles-co/algorithm_questions/trees"
)

func TestMinDiffBST(t *testing.T) {
	tt := []struct {
		name     string
		input    *trees.TreeNode
		expected int
	}{
		{
			name: "Example 1",
			input: &trees.TreeNode{
				Val: 4,
				Left: &trees.TreeNode{
					Val: 2,
					Left: &trees.TreeNode{
						Val: 1,
					},
					Right: &trees.TreeNode{
						Val: 3,
					},
				},
				Right: &trees.TreeNode{
					Val: 6,
				},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			input: &trees.TreeNode{
				Val: 1,
				Right: &trees.TreeNode{
					Val: 3,
					Left: &trees.TreeNode{
						Val: 2,
					},
				},
			},
			expected: 1,
		},
		{
			name: "Example 3",
			input: &trees.TreeNode{
				Val: 236,
				Left: &trees.TreeNode{
					Val: 104,
					Right: &trees.TreeNode{
						Val: 227,
					},
				},
				Right: &trees.TreeNode{
					Val: 701,
					Right: &trees.TreeNode{
						Val: 911,
					},
				},
			},
			expected: 9,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := trees.MinDiffBST(tc.input)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}

}
