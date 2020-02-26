package main

import "fmt"

func minusPow(x float64, n int) float64 {
	powList := []float64{1.0, 1.0 / x}
	i, opt := 2, 1.0/x
	for i <= n {
		opt *= opt
		i *= 2
		powList = append(powList, opt)
	}
	curPow, res, index := 0, 1.0, len(powList)
	for curPow != n {
		index--
		i /= 2
		opt := powList[index]
		if curPow+i <= n {
			res *= opt
			curPow += i
		}
	}
	return res
}

func posPow(x float64, n int) float64 {
	powList := []float64{1.0, x}
	i, opt := 2, x
	for i <= n {
		opt *= opt
		powList = append(powList, opt)
		i *= 2
	}
	curPow, res, index := 0, 1.0, len(powList)
	for curPow != n {
		index--
		i /= 2
		opt := powList[index]
		if curPow+i <= n {
			res *= opt
			curPow += i
		}
	}
	return res
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n > 0 {
		return posPow(x, n)
	}
	return minusPow(x, -n)
}

func main() {
	fmt.Println(minusPow(2, 4))
}
