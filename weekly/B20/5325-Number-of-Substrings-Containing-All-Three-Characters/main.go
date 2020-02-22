package main

import "fmt"

var m map[byte]int

func check() bool {
	if m['a'] > 0 && m['b'] > 0 && m['c'] > 0 {
		return true
	}
	return false
}

func numberOfSubstrings(s string) int {
	m = make(map[byte]int)
	m['a'] = 0
	m['b'] = 0
	m['c'] = 0
	l, r := 0, 0
	res := 0
	for r < len(s) {
		if check() {
			res += len(s) - r + 1
			m[s[l]]--
			l++
		} else {
			m[s[r]]++
			r++
		}
	}
	for l < len(s) {
		if check() {
			res += 1
			m[s[l]]--
			l++
		} else {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
}
