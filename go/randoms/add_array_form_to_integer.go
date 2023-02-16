package randoms

func AddToArrayForm(A []int, K int) []int {
	for i := len(A) - 1; i > -1; i-- {
		A[i] += K % 10
		K /= 10
		if A[i] > 9 {
			A[i] -= 10
			K++
		}
	}
	res := A[:]

	for K > 0 {
		tmp := make([]int, len(res)+1)
		copy(tmp[1:], res)
		tmp[0] = K % 10
		K /= 10
		res = tmp
	}

	return res
}
