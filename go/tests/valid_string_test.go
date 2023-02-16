package tests

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestValidString(t *testing.T) {
	tt := []struct {
		name     string
		s        string
		expected bool
	}{
		{name: "Test 1", s: "()", expected: true},
		{name: "Test 2", s: "()[]{}", expected: true},
		{name: "Test 3", s: "(]", expected: false},
		{name: "Test 4", s: "([)]", expected: false},
		{name: "Test 5", s: "{[]}", expected: true},
		{name: "Test 6", s: "]", expected: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := randoms.ValidString(tc.s); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
