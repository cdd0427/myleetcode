package main

//func isIsomorphic(s string, t string) bool {
//	if len(s)!= len(t){
//		return false
//	}
//	s2t:=make(map[byte]byte)
//	t2s:=make(map[byte]byte)
//	for i:=0;i<len(s);i++{
//		a,b:=s[i],t[i]
//		opt1,ok1:=s2t[a]
//		opt2,ok2:=t2s[b]
//		if opt1==b{
//			continue
//		}
//		if !ok1&&!ok2{
//			s2t[a]=b
//			t2s[b]=a
//		}
//		if (ok2&&opt2!=a)||(ok1&&opt1!=b){
//			return false
//		}
//	}
//	return true
//}

func isIsomorphic(s string, t string) bool {
	sPreIndex := [255]int{}
	tPreIndex := [255]int{}
	for i := 0; i < len(s); i++ {
		if sPreIndex[int(s[i])] != tPreIndex[(t[i])] {
			return false
		} else {
			sPreIndex[int(s[i])] = i + 1
			tPreIndex[int(t[i])] = i + 1
		}
	}
	return true
}
