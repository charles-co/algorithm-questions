package tests

import (
	"github.com/charles-co/algorithm_questions/trees"
	"testing"
)

func TestWordDictionary(t *testing.T) {

	tt := []struct {
		name     string
		words    []string
		searches []string
		expected []bool
	}{
		{
			name:     "Test 1",
			words:    []string{"bad", "dad", "mad"},
			searches: []string{"pad", "bad", ".ad", "b.."},
			expected: []bool{false, true, true, true},
		},
		{
			name:     "Test 2",
			words:    []string{"a", "a"},
			searches: []string{".", "a", "aa", "a", ".a", "a."},
			expected: []bool{true, true, false, true, false, false},
		},
		{
			name:     "Test 3",
			words:    []string{"a", "ab"},
			searches: []string{".a", "a", "aa", "a", ".a", "a."},
			expected: []bool{false, true, false, true, false, true},
		},
		{
			name:     "Test 4",
			words:    []string{"a", "ab"},
			searches: []string{"a", "a.", ".a", ".b", "ab", ".ab", "a.b", "ab.", "a..", ".a.", "..b"},
			expected: []bool{true, true, false, true, true, false, false, false, false, false, false},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			root := trees.WordDictionary()
			for _, word := range tc.words {
				root.AddWord(word)
			}

			for i, search := range tc.searches {
				res := root.Search(search)
				if res != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected[i], res)
				}
			}
		})
	}

}

