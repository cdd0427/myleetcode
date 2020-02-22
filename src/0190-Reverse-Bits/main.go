package main

import "fmt"

func d2b(n uint32) string {
	res := ""
	for n > 1 {
		res += string('0' + n%2)
		n /= 2
	}
	res += string('0' + n)
	for len(res) < 32 {
		res += "0"
	}
	return res
}

func reverseBits(num uint32) uint32 {
	var (
		res uint32
		opt uint32
	)
	res, opt = 0, 1
	s := d2b(num)
	fmt.Println(s)
	for i := len(s) - 1; i >= 0; i-- {
		res += uint32(s[i]-'0') * opt
		opt *= 2
	}
	return uint32(res)
}

func main() {
	fmt.Println(reverseBits(43261596))
}
