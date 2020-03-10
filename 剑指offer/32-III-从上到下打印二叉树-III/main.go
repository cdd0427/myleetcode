package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res [][]int
)

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level >= len(res) {
		res = append(res, []int{})
	}
	if level%2 == 0 {
		res[level] = append(res[level], root.Val)
	} else {
		res[level] = append([]int{root.Val}, res[level]...)
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
	return
}

func levelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}
