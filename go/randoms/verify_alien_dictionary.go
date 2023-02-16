package randoms

import (
	"math"
)

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
