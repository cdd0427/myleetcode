package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root, rootMirror *TreeNode) bool {
	if root == nil && rootMirror == nil {
		return true
	}
	if root == nil || rootMirror == nil {
		return false
	}
	l := helper(root.Left, rootMirror.Right)
	r := helper(root.Right, rootMirror.Left)
	if root.Val != rootMirror.Val {
		return false
	}
	return l && r
}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}
