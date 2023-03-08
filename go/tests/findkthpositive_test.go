package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestFindKthPositive(t *testing.T) {
	tt := []struct {
		name     string
		arr      []int
		k        int
		expected int
	}{
		{name: "Test 1", arr: []int{2, 3, 4, 7, 11}, k: 5, expected: 9},
		{name: "Test 2", arr: []int{1, 2, 3, 4}, k: 2, expected: 6},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := randoms.FindKthPositive(tc.arr, tc.k); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
