package tests

import (
	"github.com/charles-co/algorithm_questions/randoms"
	"testing"
)

func TestAlienSorted(t *testing.T) {
	tt := []struct {
		name     string
		words    []string
		order    string
		expected bool
	}{
		{"test1", []string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{"test2", []string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{"test3", []string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
		{"test4", []string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := randoms.IsAlienSorted(tc.words, tc.order)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
