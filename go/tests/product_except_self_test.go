package tests

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "Test 1", input: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{name: "Test 2", input: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := randoms.ProductExceptSelf(tc.input)
			if len(result) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, result)
				}
			}
		})
	}
}
