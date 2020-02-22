package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	res []int = []int{}
)

func dfs(root *ListNode) {
	if root == nil {
		return
	}
	dfs(root.Next)
	res = append(res, root.Val)
}

func reversePrint(head *ListNode) []int {
	dfs(head)
	return res
}
