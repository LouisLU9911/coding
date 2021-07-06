package leetcode_cn

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	return slow
}
