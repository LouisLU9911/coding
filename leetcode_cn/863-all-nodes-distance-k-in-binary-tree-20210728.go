package leetcode_cn

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	ret := &[]int{}
	upK(root, target, k, ret)
	return *ret
}

func upK(root *TreeNode, target *TreeNode, k int, res *[]int) int {
	if root == nil {
		return -1
	}
	if root == target {
		downK(root, 0, k, res)
		return 0
	}
	leftVal := upK(root.Left, target, k, res)
	rightVal := upK(root.Right, target, k, res)
	if leftVal == -1 && rightVal == -1 {
		return -1
	}
	if leftVal != -1 {
		if leftVal+1 == k {
			*res = append(*res, root.Val)
			return -1
		} else if leftVal+1 > k {
			return -1
		} else {
			downK(root.Right, leftVal+2, k, res)
			return leftVal + 1
		}
	}
	if rightVal != -1 {
		if rightVal+1 == k {
			*res = append(*res, root.Val)
			return -1
		} else if rightVal+1 > k {
			return -1
		} else {
			downK(root.Left, rightVal+2, k, res)
			return rightVal + 1
		}
	}
	return -1
}

func downK(root *TreeNode, n, k int, res *[]int) {
	if root == nil {
		return
	}
	if n == k {
		*res = append(*res, root.Val)
	} else if n < k {
		downK(root.Left, n+1, k, res)
		downK(root.Right, n+1, k, res)
	}
}
