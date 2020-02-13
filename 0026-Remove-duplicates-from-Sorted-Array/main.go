package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	cur, curIndex := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != cur {
			cur = nums[i]
			curIndex++
			nums[curIndex] = cur
		}
	}
	nums = nums[:curIndex+1]
	return curIndex + 1
}
