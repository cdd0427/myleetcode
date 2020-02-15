package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		pa = pa.Next
		pb = pb.Next
		if pa == nil && pb == nil {
			return nil
		}
		if pa == nil {
			pa = headB
		}
		if pb == nil {
			pb = headA
		}
	}
	return pa
}
