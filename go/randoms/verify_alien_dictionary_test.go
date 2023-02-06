package randoms_test

import (
	"testing"
	"math"
)

func isAlienSorted(words []string, order string) bool {
    mappings := make(map[string]int)
    for i, letter := range order {
        mappings[string(letter)] = i
    }

    for i := 1; i < len(words); i++ {
        first := words[i - 1]
        second := words[i]
        flag := false
        n := math.Min(float64(len(first)), float64(len(second)))

        for j := 0; j < int(n); j++ {
            if mappings[string(first[j])] > mappings[string(second[j])] {
                return false
            }
            if mappings[string(first[j])] < mappings[string(second[j])] {
                flag = true
                break;
            }
        }
        if flag == false && len(first) > len(second) {
            return false
        }
    }
    return true
}


func TestAlienSorted(t *testing.T) {
	tt := []struct{
		name string
		words []string
		order string
		expected bool
	}{
		{"test1", []string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{"test2", []string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz", false},
		{"test3", []string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz", false},
		{"test4", []string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz", false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := isAlienSorted(tc.words, tc.order)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}