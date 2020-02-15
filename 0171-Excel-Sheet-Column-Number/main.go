package main

import "fmt"

func titleToNumber(s string) int {
	b := []byte(s)
	res := 0
	for i := 0; i < len(b); i++ {
		res += res*25 + int(b[i]-64)
	}
	return res
}

func main() {
	fmt.Println(titleToNumber("ZY"))
}
