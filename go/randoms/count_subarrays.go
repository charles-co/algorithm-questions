package randoms

import (
	"math"
)

func CountSubarrays(A []int, maxK, minK int) int {
	jmin, jmax, jbad := -1, -1, -1
	res := 0

	for i := 0; i < len(A); i++ {
		if A[i] < minK || A[i] > maxK {
			jbad = i
		}

		if A[i] == minK {
			jmin = i
		}

		if A[i] == maxK {
			jmax = i
		}

		start := int(math.Min(float64(jmin), float64(jmax)))

		if jbad < start {
			res += start - jbad
		}
	}

	return res
}
