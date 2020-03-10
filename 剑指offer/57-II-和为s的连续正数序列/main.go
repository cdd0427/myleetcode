package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	if target == 1 {
		return nil
	}
	res := [][]int{}
	l, r, sum := 1, 1, 1
	for l <= r && r <= (target+1)/2 {
		if sum < target {
			r++
			sum += r
		} else if sum > target {
			sum -= l
			l++
		} else {
			opt := make([]int, r-l+1)
			for i := 0; i < len(opt); i++ {
				opt[i] = i + l
			}
			res = append(res, opt)
			sum -= l
			l++
		}
	}
	return res
}

func main() {
	fmt.Println(findContinuousSequence(15))
}
