package main

import "fmt"

//不保证顺序
func exchange(nums []int) []int {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			nums[c], nums[i] = nums[i], nums[c]
			c++
		}
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
