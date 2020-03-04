package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return dfs(a.Left, b.Left) && dfs(a.Right, b.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if B == nil || A == nil {
		return false
	}
	var res bool
	if A.Val == B.Val {
		res = dfs(A, B)
	}
	if !res {
		res = isSubStructure(A.Left, B)
	}
	if !res {
		res = isSubStructure(A.Right, B)
	}
	return res
}
