package main

import "fmt"

func integerBreak(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]
		for j := 2; i-j >= 1; j++ {
			opt := dp[i-j] * j
			if opt > dp[i] {
				dp[i] = opt
			}
		}
		if i > dp[i] && i < n {
			dp[i] = i
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(10))
}
