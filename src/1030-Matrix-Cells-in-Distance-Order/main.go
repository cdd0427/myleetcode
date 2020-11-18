package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

//解法一
//直接构造所有的点，按距离排序
func allCellsDistOrder1(R int, C int, r0 int, c0 int) [][]int {
	res := [][]int{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		d1 := abs(res[i][0]-r0) + abs(res[i][1]-c0)
		d2 := abs(res[j][0]-r0) + abs(res[j][1]-c0)
		return d1 < d2
	})
	return res
}

//解法2 BFS遍历
func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	visit := make([][]bool, R)
	for i := range visit {
		visit[i] = make([]bool, C)
	}
	//fmt.Println(visit)
	queue := [][]int{[]int{r0, c0}}
	res := [][]int{}
	step := func(i int, j int) {
		if i >= 0 && j >= 0 && i < R && j < C && !visit[i][j] {
			if !visit[i][j] {
				queue = append(queue, []int{i, j})
				visit[i][j] = true
			}
		}
	}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		res = append(res, p)
		for _, dir := range dirs {
			i, j := p[0]+dir[0], p[1]+dir[1]
			step(i, j)
		}
	}
	return res
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	visit := make([][]bool, R)
	for i := range visit {
		visit[i] = make([]bool, C)
	}
	visit[r0][c0] = true
	queue := [][]int{[]int{r0, c0}}
	res := [][]int{}
	step := func(i int, j int) {
		if i >= 0 && j >= 0 && i < R && j < C && !visit[i][j] {
			queue = append(queue, []int{i, j})
			visit[i][j] = true
		}
	}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		res = append(res, p)
		for _, dir := range dirs {
			i, j := p[0]+dir[0], p[1]+dir[1]
			step(i, j)
		}
	}
	return res
}

func main() {
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
}
