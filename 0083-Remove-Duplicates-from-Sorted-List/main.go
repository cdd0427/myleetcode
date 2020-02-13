package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil {
		op := p.Val
		pp := p
		for pp != nil && p.Val == op {
			pp = pp.Next
		}
		p.Next = pp
		p = p.Next
	}
	return head
}
