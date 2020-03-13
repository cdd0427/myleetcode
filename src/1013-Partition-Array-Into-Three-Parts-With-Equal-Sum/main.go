package main

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}
	dp := make([]int, len(A))
	dp[0] = A[0]
	for i := 1; i < len(A); i++ {
		dp[i] = dp[i-1] + A[i]
	}
	sum := dp[len(dp)-1]
	if sum%3 != 0 {
		return false
	}
	num1, num2 := sum/3, sum/3*2
	l, r := 0, len(dp)-2
	for l < r {
		if dp[l] == num1 && dp[r] == num2 {
			return true
		}
		if dp[l] != num1 {
			l++
		}
		if dp[r] != num2 {
			r--
		}
	}
	return false
}
