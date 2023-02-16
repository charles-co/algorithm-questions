package tests

import (
	"github.com/charles-co/algorithm_questions/dfs"
	"testing"
)

func TestLongestMatchingWord(t *testing.T) {
	tt := []struct {
		name       string
		dictionary []string
		chars      string
		result     string
	}{
		{"Test 1", []string{"when", "what", "whatthen", "whatnow"}, "whatno", "what"},
		{"Test 2", []string{"when", "what", "whatthen", "whatnow"}, "whatnwo", "whatnow"},
		{"Test 3", []string{"when", "what", "whatthen", "whatnow"}, "wonthaw", "whatnow"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := dfs.LongestMatchingWord(tc.dictionary, tc.chars)
			if result != tc.result {
				t.Errorf("Expected %v, got %v", tc.result, result)
			}
		})
	}

}
