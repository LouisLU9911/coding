package leetcode_cn

func pathInZigZagTree(label int) []int {
	row, leftIndex := getRowAndLeftIndex(label)
	ret := make([]int, row)
	ret[row-1] = label
	for i, thisRow := 1, row; i < row; i++ {
		leftIndex = ((leftIndex + (1 << (thisRow - 1)) - 1) / 2) - (1 << (thisRow - 2)) + 1
		thisRow--
		res := getNum(thisRow, leftIndex)
		ret[row-i-1] = res
	}
	return ret
}

func getRowAndLeftIndex(n int) (int, int) {
	row := 1
	for n > (1<<row)-1 {
		row = row + 1
	}
	isEven := (row % 2) == 0
	rest := n - (1 << (row - 1)) + 1
	var leftIndex int
	if isEven {
		// from right to left
		leftIndex = (1 << (row - 1)) - rest + 1
	} else {
		// from left to right
		leftIndex = rest
	}
	return row, leftIndex
}

func getNum(row, leftIndex int) int {
	if row%2 == 0 {
		return (1 << row) - leftIndex
	} else {
		return (1 << (row - 1)) - 1 + leftIndex
	}
}
