package randoms

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:

// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

func CheckInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	con_s1 := [26]int{}
	con_s2 := [26]int{}

	for i := 0; i < n1; i++ {
		con_s1[s1[i]-'a']++
	}

	for i := 0; i < n2; i++ {
		con_s2[s2[i]-'a']++
		if i >= n1 {
			con_s2[s2[i-n1]-'a']--
		}
		if con_s1 == con_s2 {
			return true
		}
	}

	return false

}
