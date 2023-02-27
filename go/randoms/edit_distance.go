package randoms

import (
	"math"
	"strconv"
)

// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

// You have the following three operations permitted on a word:

// Insert a character
// Delete a character
// Replace a character

func EditDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j-1]), float64(dp[i-1][j])), float64(dp[i][j-1])) + 1)
			}
		}
	}
	return dp[m][n]
}

func EditDistanceWithRecursion(word1 string, word2 string, i, j int, memo map[string]int) int {
	if i == len(word1) && j == len(word2) {
		return 0
	}
	if i == len(word1) {
		return len(word2) - j
	}
	if j == len(word2) {
		return len(word1) - i
	}
	key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
	var ans int

	if _, ok := memo[key]; !ok {
		if word1[i] == word2[j] {
			ans = EditDistanceWithRecursion(word1, word2, i+1, j+1, memo)
		} else {
			insert := EditDistanceWithRecursion(word1, word2, i, j+1, memo) + 1
			delete := EditDistanceWithRecursion(word1, word2, i+1, j, memo) + 1
			replace := EditDistanceWithRecursion(word1, word2, i+1, j+1, memo) + 1
			
			ans = int(math.Min(math.Min(float64(insert), float64(delete)), float64(replace)))

		}
		memo[key] = ans
	}
	return memo[key]
}