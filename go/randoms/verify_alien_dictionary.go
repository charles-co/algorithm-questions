package randoms

import (
	"math"
)

// In an alien language, surprisingly, they also use English lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.

// Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the given words are sorted lexicographically in this alien language.

 

// Example 1:

// Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
// Output: true
// Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
// Example 2:

// Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
// Output: false
// Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
// Example 3:

// Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
// Output: false
// Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).
 

// Constraints:

// 1 <= words.length <= 100
// 1 <= words[i].length <= 20
// order.length == 26
// All characters in words[i] and order are English lowercase letters.

func IsAlienSorted(words []string, order string) bool {
	mappings := make(map[string]int)
	for i, letter := range order {
		mappings[string(letter)] = i
	}

	for i := 1; i < len(words); i++ {
		first := words[i-1]
		second := words[i]
		flag := false
		n := math.Min(float64(len(first)), float64(len(second)))

		for j := 0; j < int(n); j++ {
			if mappings[string(first[j])] > mappings[string(second[j])] {
				return false
			}
			if mappings[string(first[j])] < mappings[string(second[j])] {
				flag = true
				break
			}
		}
		if flag == false && len(first) > len(second) {
			return false
		}
	}
	return true
}
