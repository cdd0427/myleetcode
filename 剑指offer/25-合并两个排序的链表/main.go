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
	var res *ListNode
	if l1.Val > l2.Val {
		res = l2
		l2 = l2.Next
	} else {
		res = l1
		l1 = l1.Next
	}
	head := res
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			res.Next = l2
			res = res.Next
			l2 = l2.Next
		} else {
			res.Next = l1
			res = res.Next
			l1 = l1.Next
		}
	}
	for l1 != nil {
		res.Next = l1
		l1 = l1.Next
		res = res.Next
	}
	for l2 != nil {
		res.Next = l2
		l2 = l2.Next
		res = res.Next
	}
	return head
}
