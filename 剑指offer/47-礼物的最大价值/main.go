package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		opt := make([]int, len(grid[0]))
		dp[i] = opt
	}
	var i, j int
	for i = 0; i < len(dp); i++ {
		for j = 0; j < len(dp[0]); j++ {
			up, left := 0, 0
			if i-1 >= 0 {
				up = dp[i-1][j]
			}
			if j-1 >= 0 {
				left = dp[i][j-1]
			}
			dp[i][j] = grid[i][j] + max(up, left)
		}
	}
	return dp[i-1][j-1]
}
