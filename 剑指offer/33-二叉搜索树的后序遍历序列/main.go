package main

import "fmt"

func helper(postOrder []int, l, r int) bool {
	if l >= r {
		return true
	}
	root := postOrder[r]
	i := l
	for ; i < r; i++ {
		if postOrder[i] > root {
			break
		}
	}
	for j := i; j < r; j++ {
		if postOrder[j] < root {
			return false
		}
	}
	return helper(postOrder, l, i-1) && helper(postOrder, i, r-1)
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	return helper(postorder, 0, len(postorder)-1)
}

func main() {
	fmt.Println(verifyPostorder([]int{5, 4, 3, 2, 1}))
}
