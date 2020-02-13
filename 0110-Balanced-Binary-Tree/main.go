package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	flag = true
)

func preOrder(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	l := preOrder(root.Left, level+1)
	r := preOrder(root.Left, level+1)
	if l-r > 1 || r-1 > 1 {
		flag = false
	}
	return level
}

func isBalanced(root *TreeNode) bool {
	preOrder(root, 0)
	return flag
}
