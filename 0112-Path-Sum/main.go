package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum int, target int) bool {
	if root.Left == nil && root.Right == nil {
		if sum == target {
			return true
		}
		return false
	}
	l, r := false, false
	if root.Left != nil {
		l = dfs(root.Left, sum+root.Left.Val, target)
	}
	if root.Right != nil {
		r = dfs(root.Right, sum+root.Right.Val, target)
	}
	return l || r
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, root.Val, sum)
}
