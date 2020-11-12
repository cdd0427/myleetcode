package main

import "fmt"

//双指针法原地修改
func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]&1 == 1 {
			for A[j]&1 == 1 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParityII([]int{2, 3, 4, 5}))
}
