package main

func mySqrt(x int) int {
	l, r := 0, x
	mid := (l + r) / 2
	for l <= r {
		op := mid * mid
		op2 := (mid + 1) * (mid + 1)
		if op == x || (op < x && op2 > x) {
			return mid
		}
		if op2 == x {
			return mid + 1
		}
		if op < x {
			r = mid
		}
		if op > x {
			l = mid
		}
		mid = (l + r) / 2
	}
	return mid
}
