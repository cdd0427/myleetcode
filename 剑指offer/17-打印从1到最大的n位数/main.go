package main

func printNumbers(n int) []int {
	m := 1
	for n != 0 {
		m *= 10
		n--
	}
	res := make([]int, n-1)
	for i := 1; i <= n-1; i++ {
		res[i-1] = i
	}
	return res
}
