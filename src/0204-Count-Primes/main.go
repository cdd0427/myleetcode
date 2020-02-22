package main

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	l := make([]bool, n+1)
	res := 0
	l[0], l[1] = true, true
	for i := 2; i < n; i++ {
		if !l[i] {
			res++
			for j := 2 * i; j < n; j += i {
				l[j] = true
			}
		}
	}
	return res
}
func main() {

}
