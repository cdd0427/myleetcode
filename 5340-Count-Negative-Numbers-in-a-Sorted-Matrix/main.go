package main

func countNegatives(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] < 0 {
				res++
			}
		}
	}
	return res
}
