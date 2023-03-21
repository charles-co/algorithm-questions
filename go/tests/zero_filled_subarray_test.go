package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestZeroFilledSubarray(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test 1",
			input:    []int{1, 0, 1, 1, 0},
			expected: 2,
		},
		{
			name:     "Test 2",
			input:    []int{1, 3, 0, 0, 2, 0, 0, 4},
			expected: 6,
		},
		{
			name:     "Test 3",
			input:    []int{0, 0, 0, 2, 0, 0},
			expected: 9,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := randoms.ZeroFilledSubarray(tc.input)
			if res != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, res)
			}
		})
	}
}
