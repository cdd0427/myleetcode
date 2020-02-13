package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	op, flag := digits[len(digits)-1]+1, 0
	if op >= 10 {
		op -= 10
		flag = 1
	}
	digits[len(digits)-1] = op
	for i := len(digits) - 2; i >= 0 && flag == 1; i-- {
		op = digits[i] + flag
		if op < 10 {
			digits[i] = op
			return digits
		}
		op -= 10
		digits[i] = op
	}
	if flag == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	in := []int{9, 9, 9}
	fmt.Println(plusOne(in))
}
