package tests

import (
	"github.com/charles-co/algorithm_questions/trees"
	"testing"
)

func TestIsComplete(t *testing.T) {

	tt := []struct {
		name     string
		root     *trees.TreeNode
		expected bool
	}{
		{
			name: "Test 1",
			root: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val: 2,
					Left: &trees.TreeNode{
						Val: 4,
					},
					Right: &trees.TreeNode{
						Val: 5,
					},
				},
				Right: &trees.TreeNode{
					Val: 3,
					Right: &trees.TreeNode{
						Val: 7,
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := trees.IsComplete(tc.root)

			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}

}
