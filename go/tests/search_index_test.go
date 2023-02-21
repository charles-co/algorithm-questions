package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestSearchInsert(t *testing.T){
	tt := []struct{
		name string
		input []int
		target int
		expected int
	}{
		{name: "Test 1", input: []int{1,3,5,6}, target: 5, expected: 2},
		{name: "Test 2", input: []int{1,3,5,6}, target: 2, expected: 1},
		{name: "Test 3", input: []int{1,3,5,6}, target: 7, expected: 4},
		{name: "Test 4", input: []int{1,3,5,6}, target: 0, expected: 0},
		{name: "Test 5", input: []int{1}, target: 0, expected: 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			result := randoms.SearchInsert(tc.input, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}