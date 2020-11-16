package main

import (
	"fmt"
	"sort"
)

//从低到高插入 36ms beat 14%
//func reconstructQueue(people [][]int) [][]int {
//	sort.Slice(people, func(i, j int) bool {
//		return people[i][0]<people[j][0]
//	})
//	res:=make([][]int,len(people))
//	for i:= range people{
//		k:=people[i][1]
//		for j:=range res{
//			if res[j]==nil&&k==0{
//				res[j]=people[i]
//				break
//			}
//			if res[j]==nil{
//				k--
//			}else if res[j][0]==people[i][0]{
//				k--
//			}
//
//		}
//	}
//	return res
//}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0]>people[j][0]||people[i][0]==people[j][0]&&people[i][1]<people[j][1]
	})
	for i,item:=range people{
		idx:=item[1]
		//append速度比copy慢很多
		//res = append(res[:idx],append([][]int{item},res[idx:]...)...)
		copy(people[idx+1:i+1],people[idx:i])
		people[idx]=item
	}
	return people
}



func main() {
	fmt.Println(reconstructQueue([][]int{
		{7,0},{4,4},{7,1},{5,0},{6,1},{5,2},
	}))
}
