package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	max, min := 0, 0
	if p.Val > q.Val {
		max, min = p.Val, q.Val
	} else {
		max, min = q.Val, p.Val
	}
	for root != nil {
		if root.Val >= min && root.Val <= max {
			return root
		}
		if root.Val < min {
			root = root.Left
		} else if root.Val > max {
			root = root.Right
		}
	}
	return nil
}
