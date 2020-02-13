package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = &ListNode{l1.Val, nil}
			p = p.Next
			l1 = l1.Next
		} else if l1.Val == l2.Val {
			p.Next = &ListNode{l1.Val, nil}
			p = p.Next
			p.Next = &ListNode{l2.Val, nil}
			p = p.Next
			l1, l2 = l1.Next, l2.Next
		} else {
			p.Next = &ListNode{l2.Val, nil}
			p = p.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		p.Next = &ListNode{l1.Val, nil}
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		p.Next = &ListNode{l2.Val, nil}
		p = p.Next
		l2 = l2.Next
	}
	return head.Next
}
