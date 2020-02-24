package main

import "math"

func findDivisors(num int) []int {
	d := int(math.Sqrt(float64(num)))
	for num%d != 0 {
		d--
	}
	return []int{d, num / d}
}

func closestDivisors(num int) []int {
	res1 := findDivisors(num + 1)
	res2 := findDivisors(num + 2)
	if res1[1]-res1[0] > res2[1]-res2[0] {
		return res2
	}
	return res1
}
