package main

func minArray(numbers []int) int {
	if numbers[0] < numbers[len(numbers)-1] {
		return numbers[0]
	}
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}
