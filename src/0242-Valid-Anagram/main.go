package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	v := make([]int, 26)
	for i := 0; i < len(s); i++ {
		v[s[i]-'a']++
		v[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if v[i] != 0 {
			return false
		}
	}
	return true
}
