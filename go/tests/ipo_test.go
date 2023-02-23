package tests

import (
	"testing"

	"github.com/charles-co/algorithm_questions/randoms"
)

func TestFindMaximizedCapital(t *testing.T) {
	tt := []struct {
		name     string
		k        int
		w        int
		profits  []int
		capital  []int
		expected int
	}{
		{
			name:     "Example 1",
			k:        2,
			w:        0,
			profits:  []int{1, 2, 3},
			capital:  []int{0, 1, 1},
			expected: 4,
		},
		{
			name:     "Example 2",
			k:        3,
			w:        0,
			profits:  []int{1, 2, 3},
			capital:  []int{0, 1, 2},
			expected: 6,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.FindMaximizedCapital(tc.k, tc.w, tc.profits, tc.capital)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
