package main

import "math"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := ""
	index, minLen := 0, math.MaxInt32
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	for index < minLen {
		for i := 1; i < len(strs); i++ {
			if strs[i][index] != strs[0][index] {
				return res
			}
		}
		res += string(strs[0][index])
		index++
	}
	return res
}

func main() {
	in := []string{"dog", "racecar", "car"}
	println(longestCommonPrefix(in))
}
