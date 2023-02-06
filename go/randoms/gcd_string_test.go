package randoms_test


import (
	"testing"
)

// For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).

// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcdOfStrings(str1, str2 string) string {
    if str1 + str2 != str2 + str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func TestGcdOfStrings(t *testing.T) {
	tt := []struct{
		str1 string
		str2 string
		expected string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"ABABAB", "ABAD", ""},
		{"LEET", "CODE", ""},
	}

	for _, tc := range tt {
		t.Run(tc.str1 + tc.str2, func(t *testing.T) {
			if got := gcdOfStrings(tc.str1, tc.str2); got != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, got)
			}
		})
	}
}