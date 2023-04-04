package randoms

// Given a string s, partition the string into one or more substrings such that the characters in each substring are unique. That is, no letter appears in a single substring more than once.

// Return the minimum number of substrings in such a partition.

// Note that each character should belong to exactly one substring in a partition.

// Example 1:

// Input: s = "abacaba"
// Output: 4
// Explanation:
// Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
// It can be shown that 4 is the minimum number of substrings needed.
// Example 2:

// Input: s = "ssssss"
// Output: 6
// Explanation:
// The only valid partition is ("s","s","s","s","s","s").

// Constraints:

// 1 <= s.length <= 105
// s consists of only English lowercase letters.

func PartitionString(s string) int {

	n := len(s)
	mapping := make(map[byte]int)
	i, j := 0, 1
	mapping[s[i]] = 1
	res := 1

	for i < n {
		for j < n && (mapping[s[j]] == 0 || mapping[s[j]] < i) {
			mapping[s[j]] = j
			j++
		}
		if j >= n {
			break
		}

		i, j = j, j+1
		mapping[s[i]] = i
		res++
	}
	return res
}
