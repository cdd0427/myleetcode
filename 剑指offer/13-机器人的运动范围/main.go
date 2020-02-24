package main

var (
	visited [][]bool
	res     int
)

func varInit(m, n int) {
	res = 0
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		tmp := make([]bool, n)
		visited[i] = tmp
	}
}

func dfs(i, j, m, n, k int) {
	if i < 0 || j < 0 || i > m-1 || j > m-1 || j%10+j/10+i%10+i/10 > k || visited[i][j] {
		return
	}
	visited[i][j] = true
	res++
	dfs(i-1, j, m, n, k)
	dfs(i, j-1, m, n, k)
	dfs(i+1, j, m, n, k)
	dfs(i, j+1, m, n, k)
}

func movingCount(m int, n int, k int) int {
	varInit(m, n)
	dfs(0, 0, m, n, k)
	return res
}
