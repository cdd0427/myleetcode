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

func isPalindrome(head *ListNode) bool {
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	p, i := head, 0
	for ; i < count/2; p, i = p.Next, i+1 {
		continue
	}
	r := reverseList(p)
	for i := 0; i < count/2; i++ {
		if r.Val != head.Val {
			return false
		}
		r, head = r.Next, head.Next
	}
	return true
}
