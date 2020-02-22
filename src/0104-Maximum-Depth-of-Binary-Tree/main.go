package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	max int
)

func preOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}
	level++
	if level > max {
		max = level
	}
	preOrder(root.Left, level)
	preOrder(root.Right, level)
}

func maxDepth(root *TreeNode) int {
	max = 0
	preOrder(root, 0)
	return max
}
