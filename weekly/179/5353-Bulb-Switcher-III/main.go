package main

func numTimesAllBlue(light []int) int {
	l, r := 0, light[0]
	res := 0
	if l == r-1 {
		res++
	}
	for i := 1; i < len(light); i++ {
		if light[i] > r {
			r = light[i]
		}
		l++
		if l == r-1 {
			res++
		}
	}
	return res
}
