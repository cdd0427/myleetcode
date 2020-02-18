package main

//两者前缀就是答案
func rangeBitwiseAnd(m int, n int) int {
	offset := 0
	for m != n {
		m >>= 1
		n >>= 1
		offset += 1
	}

	return m << uint32(offset)
}
