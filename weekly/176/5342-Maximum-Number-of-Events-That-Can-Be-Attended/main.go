package main

import (
	"fmt"
	"math"
)

func rmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxEvents(events [][]int) int {
	m := make(map[int]int)
	max := 0
	min := math.MaxInt32
	for i := 0; i < len(events); i++ {
		if events[i][1] > max {
			max = events[i][1]
		}
		if events[i][1] < min {
			min = events[i][1]
		}
		if _, ok := m[events[i][1]]; ok {
			if events[i][1]-events[i][0] < m[events[i][1]] {
				m[events[i][1]] = events[i][1] - events[i][0]
			}
		} else {
			m[events[i][1]] = events[i][1] - events[i][0]
		}
	}
	res := 0
	dp := make([]int, max+1)
	for i := min; i <= max; i++ {
		if last, ok := m[i]; ok {
			if last == 0 {
				dp[i] = dp[i-1] + 1
				if dp[i] > res {
					res = dp[i]
				}
				continue
			}
			dp[i] = rmax(dp[i-1], dp[i-last]+1)
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	fmt.Println(dp)
	return res
}

//func main() {
//	events:=[][]int{{1,2},{2,3},{3,4},{1,2}}
//	fmt.Println(maxEvents(events))
//}
