package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	if head.Next==nil{
		return head
	}
	odd,even,evenHead:=head,head.Next,head.Next
	for even!=nil&&even.Next!=nil{
		odd.Next=even.Next
		odd=odd.Next
		even.Next=odd.Next
		even=even.Next
	}
	odd.Next = evenHead
	return head
}

