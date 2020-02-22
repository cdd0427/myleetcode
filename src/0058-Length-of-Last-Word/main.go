package main

func lengthOfLastWord(s string) int {
	res, i := 0, len(s)-1
	for ; s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		res++
	}
	return res
}

func main() {
	println(lengthOfLastWord("Hello World    "))
}
