package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestMinimumTime(t *testing.T) {
	tt := []struct {
		name       string
		time       []int
		totalTrips int
		expected   int
	}{
		{name: "Test 1", time: []int{1, 2, 3}, totalTrips: 5, expected: 3},
		{name: "Test 2", time: []int{2}, totalTrips: 1, expected: 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := randoms.MinimumTime(tc.time, tc.totalTrips); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
