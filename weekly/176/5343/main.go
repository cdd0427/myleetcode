package main

import (
	"fmt"
	"sort"
)

func isPossible(target []int) bool {
	sort.Ints(target)
	fmt.Println(target)
	sum := len(target)
	count := 0
	for count != len(target) {
		fmt.Println(count, " ", sum)
		if sum != target[count] {
			for sum < target[count] {
				sum = 2*sum - 1
			}
			if sum > target[count] {
				return false
			}
		}
		count++
		sum = 2*sum - 1
	}
	return true
}

func main() {
	fmt.Println(isPossible([]int{8, 5}))
}
