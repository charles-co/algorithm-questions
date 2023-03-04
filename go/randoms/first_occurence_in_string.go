package randoms

func StrStr(haystack string, needle string) int {
	n := len(needle)
	h := len(haystack)

	if n > h {
		return -1
	}

	j := 0
	for i := 0; i < h; i++ {
		j = 0
		if haystack[i] == needle[j] && i+n-1 < h && haystack[i+n-1] == needle[n-1] {
			for j < n {
				if haystack[i+j] != needle[j] {
					break
				}
				j++
			}
		}

		if j == n {
			return i
		}
	}
	return -1
}
