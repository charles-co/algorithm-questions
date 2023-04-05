package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestOptimalString(t *testing.T) {

	tt := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test1",
			input:    "abacaba",
			expected: 4,
		},
		{
			name:     "test2",
			input:    "ssssss",
			expected: 6,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.PartitionString(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

