package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var rear *ListNode
	p := head
	for p != nil {
		opt := p
		p = p.Next
		opt.Next = rear
		rear = opt
	}
	return rear
}
