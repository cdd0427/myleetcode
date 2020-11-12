package main

import (
	"fmt"
)

var (
	res         []string
	digitsBytes []byte
	mapping     = [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
)

func visit(i int, cur []byte) {
	if i == len(digitsBytes) {
		res = append(res, string(cur))
		return
	}
	for _, item := range mapping[digitsBytes[i]-'2'] {
		tmp := make([]byte, len(cur))
		copy(tmp, cur)
		tmp = append(tmp, item)
		visit(i+1, tmp)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = []string{}
	digitsBytes = []byte(digits)
	visit(0, []byte{})
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("23456"))
}
