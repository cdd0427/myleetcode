package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	lsub := buildTree(nums, l, mid-1)
	rsub := buildTree(nums, mid+1, r)
	return &TreeNode{nums[mid], lsub, rsub}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}
