package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}else if length==1{
		return lists[0]
	}
	mid:=length/2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return merge2Lists(left,right)
}

func merge2Lists(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<l2.Val{
		l1.Next=merge2Lists(l1.Next,l2)
		return l1
	}
	l2.Next = merge2Lists(l2.Next,l1)
	return l2
}
