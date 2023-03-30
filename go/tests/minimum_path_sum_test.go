package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/randoms"
)

func TestMininumPathSum(t *testing.T){
	tt := []struct{
		name string
		input [][]int
		expected int
	}{
		{
			name: "Test 1",
			input: [][]int{
				{1,3,1},
				{1,5,1},
				{4,2,1},
			},
			expected: 7,
		},
		{
			name: "Test 2",
			input: [][]int{
				{1,2,3},
				{4,5,6},
			},
			expected: 12,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			actual := randoms.MinimumPathSum(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}	
}