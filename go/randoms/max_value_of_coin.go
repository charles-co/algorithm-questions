package randoms

import (
	"math"
)

func MaxValueofCoins(piles [][]int, k int) int {
	
	dp := make([][]int, len(piles))

	for i := range dp {
		dp[i] = make([]int, k + 1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(i int, coins int) int 

	dfs = func(i int, coins int) int {

		if i >= len(piles) || coins == 0 {
			return 0
		}

		if dp[i][coins] != -1 {
			return dp[i][coins]
		}

		dp[i][coins] = dfs(i + 1, coins)
		curPos := 0

		for j := 0; j < int(math.Min(float64(coins), float64(len(piles[i])))); j++ {
			curPos += piles[i][j]
			dp[i][coins] = int(math.Max(float64((dp[i][coins])), float64(curPos + dfs(i + 1, coins - j - 1))))
		}
		return dp[i][coins]


	}

	return dfs(0, k)
}