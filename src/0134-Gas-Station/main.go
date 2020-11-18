package main

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}
	total := 0
	for i := range cost {
		cost[i] = gas[i] - cost[i]
		total += cost[i]
	}
	if total < 0 {
		return -1
	}
	res, remain := 0, 0
	for i := range cost {
		remain += cost[i]
		if remain < 0 {
			remain = 0
			res = i + 1
		}
	}
	return res
}
