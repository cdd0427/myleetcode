package main

func romanToInt(s string) int {
	res := 0
	hash := [26]int{}
	hash['I'-'A'], hash['V'-'A'], hash['X'-'A'], hash['L'-'A'], hash['C'-'A'], hash['D'-'A'], hash['M'-'A'] = 1, 5, 10, 50, 100, 500, 1000
	for i := len(s) - 1; i >= 0; i-- {
		if i < len(s)-1 && hash[s[i]-'A'] < hash[s[i+1]-'A'] {
			res -= hash[s[i]-'A']
		} else {
			res += hash[s[i]-'A']
		}
	}
	return res
}

func main() {
	println(romanToInt("MCMXCIV"))
}
