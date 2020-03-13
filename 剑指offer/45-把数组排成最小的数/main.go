package main

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	s := make([]string, len(nums))
	for i := range nums {
		s[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i]+s[j] < s[j]+s[i]
	})
	res := ""
	for _, v := range s {
		res += v
	}
	return res
}
