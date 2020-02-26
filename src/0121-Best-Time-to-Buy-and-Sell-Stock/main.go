package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
