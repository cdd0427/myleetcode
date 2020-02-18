package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	maxLevel int = 0
	res      []int
)

func postOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == maxLevel {
		res = append(res, root.Val)
		maxLevel++
	}
	postOrder(root.Right, level+1)
	postOrder(root.Left, level+1)
}

func rightSideView(root *TreeNode) []int {
	res = []int{}
	postOrder(root, maxLevel)
	return res
}
