package tests

import (
	"testing"
	"github.com/charles-co/algorithm_questions/trees"
)
func TestMinReorder(t *testing.T){
	tt := []struct{
		name string
		connection [][]int
		n int
		expected int
	}{
		{
			name: "Test 1",
			connection: [][]int{{0,1},{1,3},{2,3},{4,0},{4,5}},
			n: 6,
			expected: 3,
		},
		{
			name: "Test 2",
			connection: [][]int{{1,0},{1,2},{3,2},{3,4}},
			n: 5,
			expected: 2,
		},
		{
			name: "Test 3",
			connection: [][]int{{1,0},{2,0}},
			n: 3,
			expected: 0,
		},
		{
			name: "Test 4",
			connection: [][]int{{1,2},{2,3},{3,4},{4,1},{1,5}},
			n: 6,
			expected: 3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			actual := trees.MinReorder(tc.n, tc.connection)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}