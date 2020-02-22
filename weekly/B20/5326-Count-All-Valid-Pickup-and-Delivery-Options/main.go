package main

import "fmt"

//一共有2n天，对每天都是送或收
//按随意顺序送，收只需要在送完之后
//dp[i]=dp[i-1]*(1+2+...+2*i-1)

func countOrders(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] * (2 * i * (2*i - 1)) / 2
	}
	return dp[n]
}

func main() {
	fmt.Println(countOrders(3))
}
