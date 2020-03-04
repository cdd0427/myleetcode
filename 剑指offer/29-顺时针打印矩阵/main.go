package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}
	direction := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0
	res := []int{}
	up, down := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1
	i, j := 0, 0
	for up <= down && l <= r {
		res = append(res, matrix[i][j])
		if i == up && j == r && d == 0 {
			d = 1
			up++
		} else if i == down && j == r && d == 1 {
			d = 2
			r--
		} else if i == down && j == l && d == 2 {
			d = 3
			down--
		} else if i == up && j == l && d == 3 {
			d = 0
			l++
		}
		i += direction[d][0]
		j += direction[d][1]
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
