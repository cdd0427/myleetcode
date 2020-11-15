package main

import "strings"

//对于两个相同长度的数字序列，最左边不同的数字决定了两个数字的大小
//要使得剩下的数字最小，即保证靠前的数字尽可能小

//最小栈

func removeKdigits(num string, k int) string {
	stack:=[]byte{}
	for i :=range num{
		digit:=num[i]
		for k>0&&len(stack)>0&&digit<stack[len(stack)-1]{
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack,digit)
	}
	stack = stack[:len(stack)-k]
	ans:=strings.TrimLeft(string(stack),"0")
	if ans == ""{
		return "0"
	}
	return ans
}