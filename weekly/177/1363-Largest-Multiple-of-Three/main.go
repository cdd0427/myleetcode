package main

var (
	cnt [10]int
	sum int
	res string
	m   map[int]int
)

func varInit() {
	cnt = [10]int{}
	sum = 0
	res = ""
	m = make(map[int]int)
	m[0] = 0
}

func del(m int) bool {
	for i := m; i <= 9; i += 3 {
		if cnt[i] != 0 {
			cnt[i]--
			return true
		}
	}
	return false
}

func largestMultipleOfThree(digits []int) string {
	varInit()
	for i := 0; i < len(digits); i++ {
		cnt[digits[i]]++
		sum += digits[i]
	}
	if sum%3 == 1 {
		if !del(1) {
			del(2)
			del(2)
		}
	}
	if sum%3 == 2 {
		if !del(2) {
			del(1)
			del(1)
		}
	}
	a, b := cnt[0], -cnt[0]
	for i := 9; i >= 0; i-- {
		b += cnt[i]
		for cnt[i] != 0 {
			res += string(i + '0')
			cnt[i]--
		}
	}
	if a != 0 && b == 0 {
		return "0"
	}
	return res
}
