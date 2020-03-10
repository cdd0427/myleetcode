package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "5352-Generate-a-String-With-Characters-That-Have-Odd-Counts"
	}
	s := countAndSay(n - 1)
	res := ""
	for i := 0; i < len(s); i++ {
		cNum, cCount := s[i], 1
		for i+1 < len(s) && s[i+1] == cNum {
			i++
			cCount++
		}
		sCount := strconv.Itoa(cCount)
		res += sCount
		res += string(cNum)
	}
	return res
}

func main() {
	println(countAndSay(30))
}
