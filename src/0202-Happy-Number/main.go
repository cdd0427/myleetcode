package main

func happy(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

func isHappy(n int) bool {
	fast, slow := n, n
	for {
		slow = happy(slow)
		fast = happy(fast)
		fast = happy(fast)
		if fast == 1 {
			return true
		}
		if fast == slow {
			break
		}
	}
	return false
}
