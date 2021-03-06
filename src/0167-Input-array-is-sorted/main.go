package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum < target {
			l++
		}
		if sum > target {
			r--
		}
	}
}
