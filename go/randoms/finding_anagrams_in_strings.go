package randoms

func FindAnagrams(s string, p string) (res []int) {
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
