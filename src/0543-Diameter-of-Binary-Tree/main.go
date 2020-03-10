package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	m := max(l, r)
	if l+r > res {
		res = l + r
	}
	return m + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}
