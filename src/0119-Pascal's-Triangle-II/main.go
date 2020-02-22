package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j >= 1; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
