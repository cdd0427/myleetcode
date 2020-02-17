package main

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}
	k = k % len(nums)
	reverse(nums[0 : len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
	return
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
