package randoms

import (
	"golang.org/x/exp/slices"
)

func TotalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	res := make([][]int, 0)
	var dfs func(queens []int, xyDiff []int, xySum []int)

	dfs = func(queens []int, xyDiff []int, xySum []int) {
		row := len(queens)

		if row == n {
			res = append(res, queens)
			return
		}

		col := 0

		for col < n {
			if !slices.Contains(queens, col) && !slices.Contains(xyDiff, row-col) && !slices.Contains(xySum, row+col) {
				dfs(append(queens, col), append(xyDiff, row-col), append(xySum, row+col))
			}
			col++
		}

	}

	dfs([]int{}, []int{}, []int{})

	return len(res)
}