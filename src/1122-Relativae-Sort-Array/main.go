package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank:=map[int]int{}
	for index,num := range arr2{
		rank[num]=index
	}
	sort.Slice(arr1,func(i,j int)bool{
		x,y :=arr1[i],arr1[j]
		xi,xo:=rank[x]
		yi,yo:=rank[y]
		if xo&&yo{
			return xi<yi
		}
		if xo||yo{
			return xo
		}
		return x<y
	})
	return arr1
}

