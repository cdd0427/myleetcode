package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	i := 0
	for ; i <= len(haystack)-len(needle); i++ {
		flag := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag == true {
			return i
		}
	}
	return -1
}

func main() {
	h := "jklas"
	n := "klas"
	println(strStr(h, n))
}
