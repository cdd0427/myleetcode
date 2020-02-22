package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res []string
	tmp []int
)

func traversal(root *TreeNode) {
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil {
		s := ""
		for i := 0; i < len(tmp)-1; i++ {
			s += strconv.Itoa(tmp[i]) + "->"
		}
		s += strconv.Itoa(tmp[len(tmp)-1])
		res = append(res, s)
	}
	if root.Left != nil {
		traversal(root.Left)
		tmp = tmp[:len(tmp)-1]
	}
	if root.Right != nil {
		traversal(root.Right)
		tmp = tmp[:len(tmp)-1]
	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res = []string{}
	traversal(root)
	return res
}
