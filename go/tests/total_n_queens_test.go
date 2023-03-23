package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/randoms"
)
func TestTotalNQueens(t *testing.T){
	tt := []struct{
		name string
		input int
		expected int
	}{
		{
			name: "Test 1",
			input: 4,
			expected: 2,
		},
		{
			name: "Test 2",
			input: 8,
			expected: 92,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			actual := randoms.TotalNQueens(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}