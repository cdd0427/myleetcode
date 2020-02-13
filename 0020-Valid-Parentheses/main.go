package main

//栈能排空就是对的

type stack []byte

func isValid(s string) bool {
	st := stack{}
	b := []byte(s)
	for _, v := range b {
		if v == '(' || v == '[' || v == '{' {
			st = append(st, v)
		}
		if v == ')' {
			if len(st) == 0 {
				return false
			}
			if st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		}
		if v == ']' {
			if len(st) == 0 {
				return false
			}
			if st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st)-1]
		}
		if v == '}' {
			if len(st) == 0 {
				return false
			}
			if st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	if len(st) != 0 {
		return false
	}
	return true
}

func main() {
	in := "{{[]}"
	println(isValid(in))
}
