package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	r := []rune(s)
	for i := 0; i < len(s)-1-i; i++ {
		if r[i] != r[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	println(isPalindrome(135232))
}
