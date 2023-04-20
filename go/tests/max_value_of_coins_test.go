package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestMaxValueOfCoins(t *testing.T) {
	tt := []struct {
		name     string
		input    [][]int
		k        int
		expected int
	}{
		{
			name: "test1",
			input: [][]int{
				{1, 100, 3},
				{7, 8, 9},
			},
			k:        2,
			expected: 101,
		},
		{
			name: "test2",
			input: [][]int{
				{100},
				{100},
				{100},
				{100},
				{100},
				{100},
				{1,1,1,1,1,1,700},
			},
			k:        7,
			expected: 706,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := randoms.MaxValueofCoins(tc.input, tc.k)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
