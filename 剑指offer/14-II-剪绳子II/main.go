package main

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	mod := 1000000007
	res := 1
	for n > 4 {
		res *= 3
		res %= mod
		n -= 3
	}
	return res * n % mod
}
