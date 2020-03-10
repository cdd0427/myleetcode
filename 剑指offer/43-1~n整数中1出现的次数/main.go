package main

func countDigitOne(n int) int {
	dp := make([]int, 11)
	for i, j := 0, 1; i < 10; i, j = i+1, j*10 {
		dp[i] = (i + 1) * j
	}
	m, res, n := n%10, 0, n/10
	if m != 0 {
		res = 1
	}
	i, j := 0, 10
	for n != 0 {
		mod := n % 10
		if mod <= 1 {
			res += mod * (m + 1 + dp[i])
		} else {
			res += j + mod*dp[i]
		}
		m = mod*j + m
		j *= 10
		i++
		n /= 10
	}
	return res
}
