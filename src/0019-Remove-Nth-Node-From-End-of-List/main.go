package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next==nil{
		return head.Next
	}
	counter,p1,p2 := 0,head,head
	for p1!=nil{
		counter++
		p1 = p1.Next
	}
	if counter==n{
		return head.Next
	}
	counter-=n
	for counter != 1{
		p2=p2.Next
		counter--
	}
	p2.Next = p2.Next.Next
	return head
}