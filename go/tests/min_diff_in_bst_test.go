package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/dfs"
)

func TestMinDiffBST(t *testing.T) {
	tt := []struct {
		name string
		input *dfs.TreeNode
		expected int
	}{
		{
			name: "Example 1",
			input: &dfs.TreeNode{
					Val: 4,
				Left: &dfs.TreeNode{
					Val: 2,
					Left: &dfs.TreeNode{
						Val: 1,
					},
					Right: &dfs.TreeNode{
						Val: 3,
					},
				},
				Right: &dfs.TreeNode{
					Val: 6,
				},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			input: &dfs.TreeNode{
				Val: 1,
				Right: &dfs.TreeNode{
					Val: 3,
					Left: &dfs.TreeNode{
						Val: 2,
					},
				},
			},
			expected: 1,
		},
		{
			name: "Example 3",
			input: &dfs.TreeNode{
				Val: 236,
				Left: &dfs.TreeNode{
					Val: 104,
					Right: &dfs.TreeNode{
						Val: 227,
					},
				},
				Right: &dfs.TreeNode{
					Val: 701,
					Right: &dfs.TreeNode{
						Val: 911,
					},
				},
			},
			expected: 9,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := dfs.MinDiffBST(tc.input)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}

}