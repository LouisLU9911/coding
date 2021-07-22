package leetcode_cn

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	rand := []*Node{}
	// use to record old list
	p2i := make(map[*Node]int)
	// use to record new list
	i2p := make(map[int]*Node)
	newHead := &Node{}
	for p, q, i := head, newHead, 0; p != nil; {
		q.Next = &Node{Val: p.Val}
		i2p[i] = q.Next
		p2i[p] = i
		rand = append(rand, p.Random)
		p = p.Next
		q = q.Next
		i++
	}
	r := newHead.Next
	for _, pointer := range rand {
		if pointer == nil {
			r.Random = nil
		} else {
			r.Random = i2p[p2i[pointer]]
		}
		r = r.Next
	}
	return newHead.Next
}
