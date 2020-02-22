package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	flag bool
)

func preOrder(q *TreeNode, p *TreeNode) {
	if q == nil && p == nil {
		return
	}
	if (q == nil && p != nil) || (p == nil && q != nil) || (q.Val != p.Val) {
		flag = false
		return
	}
	preOrder(q.Left, p.Left)
	preOrder(q.Right, p.Right)
}

func inOrder(q *TreeNode, p *TreeNode) {
	if q == nil && p == nil {
		return
	}
	if (q == nil && p != nil) || (p == nil && q != nil) {
		flag = false
		return
	}
	preOrder(q.Left, p.Left)
	if q.Val != p.Val {
		flag = false
		return
	}
	preOrder(q.Right, p.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	preOrder(p, q)
	inOrder(p, q)
	return flag
}
