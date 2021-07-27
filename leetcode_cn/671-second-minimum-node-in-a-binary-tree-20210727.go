package leetcode_cn

import "sort"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	res, has := get2thMinVal(root)
	if has {
		return res
	} else {
		return -1
	}
}

func get2thMinVal(root *TreeNode) (int, bool) {
	if root == nil {
		return -1, false
	} else if root.Left == nil {
		return root.Val, false
	} else {
		nodeVal := []int{root.Val, root.Left.Val, root.Right.Val}
		leftVal, has := get2thMinVal(root.Left)
		if has {
			nodeVal = append(nodeVal, leftVal)
		}
		rightVal, has := get2thMinVal(root.Right)
		if has {
			nodeVal = append(nodeVal, rightVal)
		}
		sort.Ints(nodeVal)
		if nodeVal[0] == nodeVal[len(nodeVal)-1] {
			return nodeVal[0], false
		} else {
			for i := 1; i < len(nodeVal); i++ {
				if nodeVal[i] != nodeVal[0] {
					return nodeVal[i], true
				}
			}
		}
	}
	return -1, false
}
