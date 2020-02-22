package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}
	cur := 1
	for cur < numRows-1 {
		row := []int{1}
		for i := 0; i < len(res[cur])-1; i++ {
			row = append(row, res[cur][i]+res[cur][i+1])
		}
		row = append(row, 1)
		res = append(res, row)
		cur++
	}
	return res
}
