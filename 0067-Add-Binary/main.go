package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) == 0 && len(b) == 0 {
		return "0"
	}
	if len(a) == 0 || a == "0" {
		return b
	}
	if len(b) == 0 || b == "0" {
		return a
	}
	ai, bi := len(a)-1, len(b)-1
	var flag uint8
	flag = 0
	res := ""
	for ; ai >= 0 && bi >= 0; ai, bi = ai-1, bi-1 {
		op := a[ai] - '0' + b[bi] + flag
		if op < '2' {
			flag = 0
			res = string(op) + res
		} else if op == '2' {
			flag = 1
			res = "0" + res
		} else if op == '3' {
			flag = 1
			res = "1" + res
		}
	}
	for ; ai >= 0; ai-- {
		op := a[ai] + flag
		if op < '2' {
			flag = 0
			res = string(op) + res
		} else if op == '2' {
			flag = 1
			res = "0" + res
		} else if op == '3' {
			flag = 1
			res = "1" + res
		}
	}
	for ; bi >= 0; bi-- {
		op := b[bi] + flag
		if op < '2' {
			flag = 0
			res = string(op) + res
		} else if op == '2' {
			flag = 1
			res = "0" + res
		} else if op == '3' {
			flag = 1
			res = "1" + res
		}
	}
	if flag == 1 {
		res = "1" + res
	}
	return res
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
