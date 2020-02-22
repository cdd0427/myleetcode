package main

func findRepeatNumber(nums []int) int {
	v := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if v[nums[i]] != 0 {
			return nums[i]
		} else {
			v[nums[i]] = 1
		}
	}
	return 0
}
