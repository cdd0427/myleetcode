package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func hasCycle(head *ListNode) bool {
//	visited:=map[*ListNode]bool{}
//	p:=head
//	for p!=nil{
//		if _,ok:=visited[p];ok{return true}
//		visited[p]=true
//		p=p.Next
//	}
//	return false
//}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
