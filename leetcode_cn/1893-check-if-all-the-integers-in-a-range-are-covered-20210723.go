package leetcode_cn

func isCovered(ranges [][]int, left int, right int) bool {
	remain := [][]int{[]int{left, right}}
	flag := false
	for i := 0; i < len(ranges); i++ {
		flag = false
		for _, this := range remain {
			if this[0] == -1 {
				continue
			} else {
				flag = true
				if ranges[i][1] < this[0] || ranges[i][0] > this[1] {
					continue
				} else if ranges[i][0] <= this[0] {
					if ranges[i][1] >= this[1] {
						this[0] = -1
					} else {
						this[0] = ranges[i][1] + 1
					}
				} else {
					if ranges[i][1] >= this[1] {
						this[1] = ranges[i][0] - 1
					} else {
						r := this[1]
						this[1] = ranges[i][0] - 1
						remain = append(remain, []int{ranges[i][1] + 1, r})
					}
				}
			}
		}
		if !flag {
			return true
		}
	}
	for i := 0; i < len(remain); i++ {
		if remain[i][0] != -1 {
			return false
		}
	}
	return true
}
