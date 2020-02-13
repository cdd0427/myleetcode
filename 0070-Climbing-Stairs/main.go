package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp1, dp2 := 1, 2
	for i := 3; i <= n; i++ {
		t := dp1 + dp2
		dp1 = dp2
		dp2 = t
	}
	return dp2
}
