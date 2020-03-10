package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res [][]int
	tmp []int
)

func dfs(root *TreeNode, sum, target int) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	sum += root.Val
	if root.Left == nil && root.Right == nil && sum == target {
		opt := make([]int, len(tmp))
		copy(opt, tmp)
		res = append(res, opt)
	}
	if root.Left != nil {
		dfs(root.Left, sum, target)
		tmp = tmp[:len(tmp)-1]
	}
	if root.Right != nil {
		dfs(root.Right, sum, target)
		tmp = tmp[:len(tmp)-1]
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	res = [][]int{}
	tmp = []int{}
	dfs(root, 0, sum)
	return res
}
