package main

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	p := &ListNode{0,head}
	res := p
	for p.Next!=nil{
		if p.Next.Next==nil{
			break
		}
		a,b:=p.Next,p.Next.Next
		p.Next,a.Next,b.Next=b,b.Next,a
		p = p.Next.Next
	}
	return res.Next
}
