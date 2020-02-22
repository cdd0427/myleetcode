package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findIndex(n int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == n {
			return i
		}
	}
	return 0
}

func build(preL, preR, inL, inR int, preorder []int, inorder []int) *TreeNode {
	if preL > preR || inL > inR {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[preL]
	m := findIndex(preorder[preL], inorder)
	root.Left = build(preL+1, preL+m-inL, inL, m-1, preorder, inorder)
	root.Right = build(preL+m-inL+1, preR, m+1, inR, preorder, inorder)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(0, len(preorder)-1, 0, len(inorder)-1, preorder, inorder)
}
