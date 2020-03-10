package main

type orange struct {
	x     int
	y     int
	round int
}

func orangesRotting(grid [][]int) int {
	fresh := 0
	res := 0
	q := []orange{}
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		opt := make([]bool, len(grid[0]))
		visited[i] = opt
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				q = append(q, orange{i, j, 0})
				visited[i][j] = true
				fresh++
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	for len(q) != 0 {
		opt := q[0]
		x, y, round := opt.x, opt.y, opt.round
		q = q[1:]
		fresh--
		grid[x][y] = 2
		if round > res {
			res = round
		}
		if x-1 >= 0 && grid[x-1][y] == 1 && visited[x-1][y] == false {
			q = append(q, orange{x - 1, y, round + 1})
			visited[x-1][y] = true
		}
		if x+1 < len(grid) && grid[x+1][y] == 1 && visited[x+1][y] == false {
			q = append(q, orange{x + 1, y, round + 1})
			visited[x+1][y] = true
		}
		if y-1 >= 0 && grid[x][y-1] == 1 && visited[x][y-1] == false {
			q = append(q, orange{x, y - 1, round + 1})
			visited[x][y-1] = true
		}
		if y+1 < len(grid[0]) && grid[x][y+1] == 1 && visited[x][y+1] == false {
			q = append(q, orange{x, y + 1, round + 1})
			visited[x][y+1] = true
		}
	}
	if fresh != 0 {
		return -1
	}
	return res
}
