package main

func explore(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	explore(grid, i-1, j)
	explore(grid, i+1, j)
	explore(grid, i, j-1)
	explore(grid, i, j+1)
}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				explore(grid, i, j)
				res++
			}
		}
	}
	return res
}
