package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	tt := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{name: "Test 1", piles: []int{3, 6, 7, 11}, h: 8, expected: 4},
		{name: "Test 2", piles: []int{30, 11, 23, 4, 20}, h: 5, expected: 30},
		{name: "Test 3", piles: []int{30, 11, 23, 4, 20}, h: 6, expected: 23},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := randoms.MinEatingSpeed(tc.piles, tc.h); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
