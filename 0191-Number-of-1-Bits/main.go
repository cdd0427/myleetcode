package main

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	var res uint32
	res = 1
	for num > 1 {
		res += num % 2
		num /= 2
	}
	return int(res)
}
