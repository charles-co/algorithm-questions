package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"test1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"test2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"test3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := randoms.TwoSum(tc.nums, tc.target)
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
