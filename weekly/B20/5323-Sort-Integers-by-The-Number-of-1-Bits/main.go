package main

import (
	"fmt"
)

func countOne(a int) int {
	c := 0
	for ; a > 0; c++ {
		a &= (a - 1)
	}
	return c
}

func sortByBits(arr []int) []int {
	v := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		v[i] = countOne(arr[i])
	}
	for i := 0; i < len(arr); i++ {
		minOne, min, minIndex := v[i], arr[i], i
		for j := i; j < len(arr); j++ {
			if v[j] < minOne {
				minOne, min, minIndex = v[j], arr[j], j
			} else if v[j] == minOne && arr[j] < min {
				minOne, min, minIndex = v[j], arr[j], j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		v[i], v[minIndex] = v[minIndex], v[i]
	}
	return arr
}

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}
