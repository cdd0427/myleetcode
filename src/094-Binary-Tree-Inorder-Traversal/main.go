package main

import "myleetcode-go/structure"

type TreeNode = structure.TreeNode

var res []int

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	res = append(res, root.Val)
	inorder(root.Right)
}

func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	inorder(root)
	return res
}
