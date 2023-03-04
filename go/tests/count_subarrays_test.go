package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestCountSubarrays(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		maxK     int
		minK     int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 3, 5, 2, 7, 5},
			maxK:     5,
			minK:     1,
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    []int{1, 1, 1, 1},
			maxK:     1,
			minK:     1,
			expected: 10,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.CountSubarrays(tc.input, tc.maxK, tc.minK)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
