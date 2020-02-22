package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func swap(root *TreeNode) {
	opt := root.Left
	root.Left = root.Right
	root.Right = opt
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = invertTree(root.Right)
	root.Left = invertTree(root.Left)
	swap(root)
	return root
}
