package main

import "fmt"

func sumSubarrayMins(A []int) int {
	s := [][2]int{}
	res, tmp := 0, 0
	for i := 0; i < len(A); i++ {
		c := 1
		for len(s) != 0 && s[len(s)-1][0] > A[i] {
			opt := s[len(s)-1]
			s = s[:len(s)-1]
			c += opt[1]
			tmp -= opt[0] * opt[1]
		}
		s = append(s, [2]int{A[i], c})
		tmp += A[i] * c
		res += tmp
		res %= 1000000007
	}
	return res
}

func main() {
	a := sumSubarrayMins([]int{3, 1, 2, 4})
	fmt.Println(a)
}
