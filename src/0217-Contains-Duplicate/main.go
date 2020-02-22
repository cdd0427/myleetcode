package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return false
		}
		m[nums[i]] = false
	}
	return true
}
