package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type treeLevel struct {
	tree  *TreeNode
	level int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := []*treeLevel{}
	queue = append(queue, &treeLevel{root, 1})
	for len(queue) != 0 {
		opt := queue[0]
		queue = queue[1:]
		if opt.tree.Left == nil && opt.tree.Right == nil {
			res = opt.level
			break
		}
		if opt.tree.Left != nil {
			queue = append(queue, &treeLevel{opt.tree.Left, opt.level + 1})
		}
		if opt.tree.Right != nil {
			queue = append(queue, &treeLevel{opt.tree.Right, opt.level + 1})
		}
	}
	return res
}
