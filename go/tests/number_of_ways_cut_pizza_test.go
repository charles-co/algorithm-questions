package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestWays(t *testing.T) {
	tt := []struct {
		name     string
		input    []string
		k        int
		expected int
	}{
		{
			name:     "Test 1",
			input:    []string{"A..", "AAA", "..."},
			k:        3,
			expected: 3,
		},
		{
			name:     "Test 2",
			input:    []string{"A..", "AA.", "..."},
			k:        3,
			expected: 1,
		},
		{
			name:     "Test 3",
			input:    []string{"A..", "A..", "..."},
			k:        1,
			expected: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.NumberOfWaysCutPizza(tc.input, tc.k)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
