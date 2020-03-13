package main

import (
	"fmt"
)

var (
	res  int
	nums []int
	set  map[string]bool
)

//携带s去重
func dfs(curIndex, prev int, s string) {
	if curIndex == len(nums) {
		fmt.Println(s + string('a'+prev))
		if _, ok := set[s]; !ok {
			set[s] = true
		}
		res++
		return
	}
	if prev == 0 {
		dfs(curIndex+1, nums[curIndex], s+string('a'))
	} else {
		if nums[curIndex]+prev*10 <= 25 {
			dfs(curIndex+1, nums[curIndex]+prev*10, s)
		}
		dfs(curIndex+1, nums[curIndex], s+string('a'+prev))
	}
}

func int2slice(num int) {
	if num != 0 {
		int2slice(num / 10)
		nums = append(nums, num%10)
	}
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	res = 0
	nums = []int{}
	int2slice(num)
	set = make(map[string]bool)
	if len(nums) == 1 {
		return 1
	}
	dfs(1, nums[0], "")
	return len(set)
}

func main() {
	fmt.Println(translateNum(1000000))
	fmt.Println(nums)
}
