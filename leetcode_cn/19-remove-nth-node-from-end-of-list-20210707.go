package leetcode_cn

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	fast, slow := head, head
	i := 1
	for ; fast != nil && i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head
	}
	if fast.Next == nil {
		return head.Next
	} else {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
