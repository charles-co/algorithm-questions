package tests

import (
	"testing"

	"github.com/charles-co/algorithm_questions/trees"
)

func TestInvertTree(t *testing.T) {
	tt := []struct {
		name     string
		input    *trees.TreeNode
		expected *trees.TreeNode
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
					Val: 7,
					Left: &trees.TreeNode{
						Val: 6,
					},
					Right: &trees.TreeNode{
						Val: 9,
					},
				},
			},
			expected: &trees.TreeNode{
				Val: 4,
				Left: &trees.TreeNode{
					Val: 7,
					Left: &trees.TreeNode{
						Val: 9,
					},
					Right: &trees.TreeNode{
						Val: 6,
					},
				},
				Right: &trees.TreeNode{
					Val: 2,
					Left: &trees.TreeNode{
						Val: 3,
					},
					Right: &trees.TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			name: "Example 2",
			input: &trees.TreeNode{
				Val: 2,
				Left: &trees.TreeNode{
					Val: 1,
				},
				Right: &trees.TreeNode{
					Val: 3,
				},
			},
			expected: &trees.TreeNode{
				Val: 2,
				Left: &trees.TreeNode{
					Val: 3,
				},

				Right: &trees.TreeNode{
					Val: 1,
				},
			},
		},
		{
			name: "Example 3",
			input: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val: 2,
				},
				Right: &trees.TreeNode{
					Val: 3,
				},
			},
			expected: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val: 3,
				},
				Right: &trees.TreeNode{
					Val: 2,
				},
			},
		},
		{
			name: "Example 4",
			input: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val: 2,
					Left: &trees.TreeNode{
						Val: 3,
					},
					Right: &trees.TreeNode{
						Val: 4,
					},
				},
				Right: &trees.TreeNode{
					Val: 5,
					Left: &trees.TreeNode{
						Val: 6,
					},
					Right: &trees.TreeNode{
						Val: 7,
					},
				},
			},
			expected: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val: 5,
					Left: &trees.TreeNode{
						Val: 7,
					},
					Right: &trees.TreeNode{
						Val: 6,
					},
				},
				Right: &trees.TreeNode{
					Val: 2,
					Left: &trees.TreeNode{
						Val: 4,
					},
					Right: &trees.TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := trees.InvertTree(tc.input)
			if !trees.IsSameTree(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
