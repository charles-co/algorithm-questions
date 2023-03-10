package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/randoms"
)

func TestLetterCombinations(t *testing.T) {
	tt := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "Example 1",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Example 2",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "Example 3",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.LetterCombinations(tc.digits)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}