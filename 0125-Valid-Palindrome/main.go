package main

import (
	"fmt"
)

func check(c rune) bool {
	if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
		return false
	}
	return true
}

func equal(a, b rune) bool {
	if ('A' <= a && a <= 'Z' && '0' <= b && b <= '9') || ('A' <= b && b <= 'Z' && '0' <= a && a <= '9') {
		return false
	}
	if a == b {
		return true
	}
	if a+32 == b {
		return true
	}
	if a == b+32 {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	b := []rune(s)
	if len(b) <= 1 {
		return true
	}
	l, r := 0, len(b)-1
	for l < r {
		for ; !check(b[l]) && l < r; l++ {
		}
		for ; !check(b[r]) && l < r; r-- {
		}
		//fmt.Println(l," ",r," ",b[l]," ",b[r])
		if !equal(b[l], b[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	in := "0P"
	fmt.Println('A' - 'a')
	fmt.Println(isPalindrome(in))
}
