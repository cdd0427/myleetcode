package main

import "fmt"

type node struct {
	x, y  int
	level int
	prev  []node
}

type queue []node

func printNode(n node) {
	fmt.Printf("节点[%d,%d] %d，其前驱为", n.x, n.y, n.level)
	for i := 0; i < len(n.prev)-1; i++ {
		fmt.Printf("  [%d,%d]", n.prev[i].x, n.prev[i].y)
	}
	fmt.Printf(" [%d,%d]\n", n.prev[len(n.prev)-1].x, n.prev[len(n.prev)-1].y)
}

func (q *queue) push(n node) {
	*q = append(*q, n)
}

func (q *queue) pop() node {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func visited(n node, x, y int) bool {
	for i := 0; i < len(n.prev); i++ {
		if x == n.prev[i].x && y == n.prev[i].y {
			return true
		}
	}
	return false
}

func bfs(byte [][]byte, word string, i, j int) bool {
	q := queue{}
	q.push(node{
		x:     i,
		y:     j,
		level: 1,
		prev:  []node{},
	})
	for len(q) != 0 {
		opt := q.pop()
		row, col, index := opt.x, opt.y, opt.level
		if index == len(word) {
			return true
		}
		if row-1 >= 0 && byte[row-1][col] == word[index] && !visited(opt, row-1, col) {
			q.push(node{row - 1, col, index + 1, append(opt.prev, opt)})
		}
		if col-1 >= 0 && byte[row][col-1] == word[index] && !visited(opt, row, col-1) {
			q.push(node{row, col - 1, index + 1, append(opt.prev, opt)})
		}
		if row+1 < len(byte) && byte[row+1][col] == word[index] && !visited(opt, row-1, col) {
			q.push(node{row + 1, col, index + 1, append(opt.prev, opt)})
		}
		if col+1 < len(byte[0]) && byte[row][col+1] == word[index] && !visited(opt, row-1, col) {
			q.push(node{row, col + 1, index + 1, append(opt.prev, opt)})
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if bfs(board, word, i, j) {
					return true
				}
			}
		}
	}
	return false
}
