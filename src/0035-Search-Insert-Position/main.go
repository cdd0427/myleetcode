package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	mid := (l + r) / 2
	for l <= r {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
			mid = (l + r) / 2
		}
		if nums[mid] < target {
			l = mid + 1
			mid = (l + r) / 2
		}
	}
	return l
}
