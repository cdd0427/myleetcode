package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	rear := head
	for rear != nil {
		opt := rear.Next
		rear.Next = prev
		prev = rear
		rear = opt
	}
	return prev
}
