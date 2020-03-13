package main

import (
	"strings"
)

//辗转相除法
func gcdOfStrings(str1 string, str2 string) string {
	s1, s2 := "", ""
	if len(str1) > len(str2) {
		s1, s2 = str1, str2
	} else {
		s1, s2 = str2, str1
	}
	for {
		if ok := strings.HasPrefix(s1, s2); ok {
			for strings.HasPrefix(s1, s2) {
				s1 = strings.TrimPrefix(s1, s2)
			}
			s1, s2 = s2, s1
		} else {
			return ""
		}
		if s2 == "" {
			return s1
		}
	}
}
