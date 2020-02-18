package main

func partitionDisjoint(A []int) int {
	min := make([]int, len(A))
	min[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] < min[i+1] {
			min[i] = A[i]
		} else {
			min[i] = min[i+1]
		}
	}
	max := A[0]
	for i := 1; i < len(A); i++ {
		if max <= min[i] {
			return i
		}
		if A[i] > max {
			max = A[i]
		}
	}
	return 0
}
