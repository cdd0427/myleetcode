package main

//func lengthOfLongestSubstring(s string) int {
//	if len(s)<=1{
//		return len(s)
//	}
//	l,r,max:=0,0,0
//	set:=make(map[byte]int)
//	for ;r< len(s);r++{
//		if i,ok:=set[s[r]];ok{
//			for ;l<=i;l++{
//				delete(set,s[i])
//			}
//		}
//		set[s[r]]=r
//		if r-l>max{
//			max=r-l
//		}
//	}
//	return max+1
//}
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	l, r, max := 0, 0, 0
	set := make(map[byte]int)
	for ; r < len(s); r++ {
		if i, ok := set[s[r]]; ok && i >= l {
			l = i + 1
		}
		set[s[r]] = r
		if r-l > max {
			max = r - l
		}
	}
	return max
}
