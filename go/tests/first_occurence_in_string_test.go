package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/randoms"
)

func TestStrSTr(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		input2   string
		expected int
	}{
		{
			name:     "Example 1",
			input:    "hello",
			input2:   "ll",
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    "aaaaa",
			input2:   "bba",
			expected: -1,
		},
		{
			name:     "Example 3",
			input:    "a",
			input2:   "a",
			expected: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.StrStr(tc.input, tc.input2)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}