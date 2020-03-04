package main

import "fmt"

//小数规则：只有一个小数点,小数点可以存在在第一位或最后一位
//1.  .1算有效数字  .不算
//指数规则：e前面可以有任意正负小数，e后面只能是正负整数
//.e不算

func isNumber(s string) bool {
	if s == "" || s == "." || s[0] == 'e' {
		return false
	}
	dotIndex, eIndex := -1, -1
	i := 0
	onlySpace := true
	for ; i < len(s) && s[i] == ' '; i++ {
		continue
	}
	end := len(s) - 1
	for ; end >= i && s[end] == ' '; end-- {
		continue
	}
	if s[0] == '+' || s[0] == '-' {
		i++
	}
	start := i
	for ; i <= end; i++ {
		if s[i] <= '9' && s[i] >= '0' {
			onlySpace = false
			continue
		} else if s[i] == '.' {
			//排除小数点在e后
			if i > eIndex && eIndex != -1 {
				return false
			}
			//排除两个小数点
			if dotIndex == -1 {
				dotIndex = i
			} else {
				return false
			}
		} else if s[i] == 'e' {
			//排除.e和e打头
			if s[i-1] == '.' || i == start {
				return false
			}
			//排除两个e
			if eIndex == -1 {
				eIndex = i
				//排除e在末尾
				if i == end {
					return false
				}
				//吸收加减
				if s[i+1] == '+' || s[i+1] == '-' {
					i++
					if i == end {
						return false
					}
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if onlySpace {
		return false
	}
	return true
}

func main() {
	fmt.Println(isNumber("+ 1"))
}
