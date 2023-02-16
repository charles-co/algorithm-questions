package tests

// Given two binary strings a and b, return their sum as a binary string.

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tt := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{name: "Test 1", a: "11", b: "1", expected: "100"},
		{name: "Test 2", a: "1010", b: "1011", expected: "10101"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := randoms.AddBinary(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
