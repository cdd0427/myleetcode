package main

var (
	m       map[int][]int
	visited []bool
	res     float64
)

func count(root int) int {
	res := 0
	for i := 0; i < len(m[root]); i++ {
		if visited[m[root][i]] == false {
			res++
		}
	}
	return res
}

func dfs(root, depth, t, target int, p float64) {
	if depth > t || res != 0 {
		return
	}
	c := count(root)
	if root == target && (t == depth || c == 0) {
		res = p
		return
	}
	for i := 0; i < len(m[root]); i++ {
		if visited[m[root][i]] == false && c != 0 {
			visited[m[root][i]] = true
			dfs(m[root][i], depth+1, t, target, p*(1.0/float64(c)))
			visited[m[root][i]] = false
		}
	}
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	m = make(map[int][]int)
	visited = make([]bool, n+1)
	res = 0
	for i := 0; i < len(edges); i++ {
		cur, next := edges[i][0], edges[i][1]
		if _, ok := m[cur]; ok {
			m[cur] = append(m[cur], next)
		} else {
			m[cur] = []int{next}
		}
		if _, ok := m[next]; ok {
			m[next] = append(m[next], cur)
		} else {
			m[next] = []int{cur}
		}
	}
	visited[1] = true
	dfs(1, 0, t, target, 1.0)
	return res
}
