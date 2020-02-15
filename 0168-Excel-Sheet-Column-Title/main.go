package main

import "fmt"

func convertToTitle(n int) string {
	res := ""
	for n != 0 {
		n -= 1
		res = string(n%26+65) + res
		n = n / 26
	}
	return res
}

func main() {
	fmt.Println(convertToTitle(52))
}
