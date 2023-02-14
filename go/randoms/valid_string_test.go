package randoms_test

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

import (
	"testing"
)

func validString(s string) bool {
	stack := []rune{}
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if r == ')' && last != '(' {
				return false
			}
			if r == '}' && last != '{' {
				return false
			}
			if r == ']' && last != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

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
			if got := validString(tc.s); got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}