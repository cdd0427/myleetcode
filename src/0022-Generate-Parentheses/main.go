package main

import "fmt"

var (
	res []string
)

func visit(remain int, gap int, cur []byte) {
	if remain == 0 {
		for gap > 0 {
			cur = append(cur, ')')
			gap--
		}
		res = append(res, string(cur))
		return
	}
	tmp := make([]byte, len(cur))
	copy(tmp, cur)
	tmp = append(tmp, '(')
	visit(remain-1, gap+1, tmp)
	tmp = make([]byte, len(cur))
	copy(tmp, cur)
	tmp = append(tmp, ')')
	if gap > 0 {
		visit(remain, gap-1, tmp)
	}
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res = []string{}
	visit(n, 0, []byte{})
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
