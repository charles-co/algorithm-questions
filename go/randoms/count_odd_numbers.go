package randoms

// Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).

func CountOdds(low int, high int) int {
	if low&1 == 0 {
		low++
	}

	if low > high {
		return 0
	} else {
		return ((high - low) / 2) + 1
	}
}
