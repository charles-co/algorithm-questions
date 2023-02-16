package tests

import (
	"testing"

	"github.com/charles-co/algorithm_questions/randoms"
	"golang.org/x/exp/slices"
)

func TestFindAnagrams(t *testing.T) {
	tt := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{name: "Test 1", s: "cbaebabacd", p: "abc", expected: []int{0, 6}},
		{name: "Test 2", s: "abab", p: "ab", expected: []int{0, 1, 2}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := randoms.FindAnagrams(tc.s, tc.p); slices.Equal(got, tc.expected) == false {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
