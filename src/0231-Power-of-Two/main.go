package main

func isPowerOfTwo(n int) bool {
	return n == n&(n-1) && n != 0
}
