package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	q := []*TreeNode{root}
	for len(q) != 0 {
		opt := q[0]
		q = q[1:]
		res = append(res, opt.Val)
		if opt.Left != nil {
			q = append(q, opt.Left)
		}
		if opt.Right != nil {
			q = append(q, opt.Right)
		}
	}
	return res
}
