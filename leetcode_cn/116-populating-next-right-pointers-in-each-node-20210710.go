package leetcode_cn

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	for l, r := root.Left, root.Right; l != nil && r != nil; {
		l.Next = r
		l = l.Right
		r = r.Left
	}
	if root.Left != nil {
		connect(root.Left)
	}
	if root.Right != nil {
		connect(root.Right)
	}
	return root
}
