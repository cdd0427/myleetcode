package main

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	i1, i2, i3 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i1]*2, dp[i2]*3, dp[i3]*5)
		if dp[i] >= dp[i1]*2 {
			i1++
		}
		if dp[i] >= dp[i2]*3 {
			i2++
		}
		if dp[i] >= dp[i3]*5 {
			i3++
		}
	}
	return dp[n]
}
