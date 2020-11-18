package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//层序遍历
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	tmp := []*Node{}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			tmp = append(tmp, cur.Left, cur.Right)
		}
		if len(queue) == 0 && len(tmp) != 0 {
			for i := 0; i < len(tmp)-1; i++ {
				tmp[i].Next = tmp[i+1]
			}
			queue = tmp
			tmp = []*Node{}
		}
	}
	return root
}
