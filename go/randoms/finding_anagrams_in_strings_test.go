package randoms_test

import (
	"golang.org/x/exp/slices"
	"testing"
)

func findAnagrams(s string, p string) (res []int) {
	n1 := len(p)
	n2 := len(s)

	con_p := [26]int{}
	con_s := [26]int{}

	for i := 0; i < n1; i++ {
		con_p[p[i]-'a']++
	}

	for i := 0; i < n2; i++ {
		con_s[s[i]-'a']++
		if i >= n1 {
			con_s[s[i-n1]-'a']--
		}
		if con_p == con_s {
			res = append(res, i-n1+1)
		}
	}

	return
}

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
			if got := findAnagrams(tc.s, tc.p); slices.Equal(got, tc.expected) == false {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
