package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	p := head
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			return head
		}
		p = p.Next
	}
	return head
}
