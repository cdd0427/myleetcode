package main

import (
	"strconv"
)

func findNum(num, index int) int {
	s := strconv.Itoa(num)
	res, _ := strconv.Atoi(string(s[index]))
	return res
}

//123 → 123-10=113
//113 → 第113/2个二位数的113%2位 → （113/2+10）/10^(113%2）
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	//10*9 2位数
	round, digits := 10, 2
	n -= round
	for {
		opt := round * 9 * digits
		if n-opt < 0 {
			break
		}
		n -= opt
		round *= 10
		digits++
	}
	return findNum(n/digits+round, n%digits)
}
