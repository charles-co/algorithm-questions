package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

// For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).

// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

func TestGcdOfStrings(t *testing.T) {
	tt := []struct {
		str1     string
		str2     string
		expected string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"ABABAB", "ABAD", ""},
		{"LEET", "CODE", ""},
	}

	for _, tc := range tt {
		t.Run(tc.str1+tc.str2, func(t *testing.T) {
			if got := randoms.GcdOfStrings(tc.str1, tc.str2); got != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, got)
			}
		})
	}
}
