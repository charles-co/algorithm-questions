package tests

// Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestCountOdds(t *testing.T) {
	tt := []struct {
		name     string
		low      int
		high     int
		expected int
	}{
		{name: "Test 1", low: 3, high: 7, expected: 3},
		{name: "Test 2", low: 8, high: 10, expected: 1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := randoms.CountOdds(tc.low, tc.high)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
