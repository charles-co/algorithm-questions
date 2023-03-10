package randoms

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letterMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var res []string
	var dfs func(int, string)
	dfs = func(idx int, s string) {
		if idx == len(digits) {
			res = append(res, s)
			return
		}

		for _, v := range letterMap[digits[idx]] {
			dfs(idx+1, s+v)
		}
	}

	dfs(0, "")
	return res
}