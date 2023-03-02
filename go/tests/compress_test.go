package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestCompress(t *testing.T) {
	tt := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "Example 1",
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
		},
		{
			name:     "Example 2",
			input:    []byte{'a'},
			expected: 1,
		},
		{
			name:     "Example 3",
			input:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: 4,
		},
		{
			name:     "Example 4",
			input:    []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			expected: 6,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.Compress(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
