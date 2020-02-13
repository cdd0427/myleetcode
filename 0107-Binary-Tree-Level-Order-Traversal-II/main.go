package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	max int
	res [][]int
)

func preOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}
	level++
	if level > max {
		max = level
		res = append(res, []int{})
	}
	res[level-1] = append(res[level-1], root.Val)
	preOrder(root.Left, level)
	preOrder(root.Right, level)
}

func reverse() {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	max = 0
	res = [][]int{}
	preOrder(root, 0)
	reverse()
	return res
}
