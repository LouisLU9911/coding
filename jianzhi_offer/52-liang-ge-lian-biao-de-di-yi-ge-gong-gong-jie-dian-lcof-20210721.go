package jianzhi_offer

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tailA := headA
	for tailA.Next != nil {
		tailA = tailA.Next
	}
	tailA.Next = headB
	fast, slow := headA, headA
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast.Next == nil {
			tailA.Next = nil
			return nil
		}
		fast = fast.Next
		if fast == slow {
			break
		}
	}
	if fast.Next == nil {
		tailA.Next = nil
		return nil
	}
	fast = headA
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	tailA.Next = nil
	return fast
}
