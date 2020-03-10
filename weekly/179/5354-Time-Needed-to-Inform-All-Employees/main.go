package main

type node struct {
	child []int
}

var res int

func dfs(tree []node, informTime []int, root int, sum int) {
	if len(tree[root].child) == 0 && sum > res {
		res = sum
	}
	for i := 0; i < len(tree[root].child); i++ {
		dfs(tree, informTime, tree[root].child[i], sum+informTime[root])
	}
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	res = 0
	tree := make([]node, n)
	for i := 0; i < n; i++ {
		if i == headID {
			continue
		}
		tree[manager[i]].child = append(tree[manager[i]].child, i)
	}
	dfs(tree, informTime, headID, 0)

	return res
}

func main() {
	numOfMinutes(7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1})
	numOfMinutes(4, 2, []int{3, 3, -1, 2}, []int{0, 0, 162, 914})
	numOfMinutes(15, 0, []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}, []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0})
	numOfMinutes(1, 0, []int{-1}, []int{0})
}
