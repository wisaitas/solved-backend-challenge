package one

import (
	"math"
)

func FindMaxPath(triangle [][]int) int {
	rows := len(triangle)

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i])
	}

	for i := rows - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] += int(math.Max(float64(dp[i+1][j]), float64(dp[i+1][j+1])))
		}
	}

	return dp[0][0]
}
