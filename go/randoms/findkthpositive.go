package randoms

func FindKthPositive(arr []int, k int) int {
	l, r := 0, len(arr)

	for l < r {
		mid := l + (r-l)/2
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l + k
}
