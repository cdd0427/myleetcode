package main

import "fmt"

func isUgly(num int) bool {
	d := 2
	for num != 1 {
		fmt.Println(num, d)
		if num%d == 0 {
			num = num / d
		} else {
			if d == 2 {
				d = 3
			} else if d == 3 {
				d = 5
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isUgly(6))
}
