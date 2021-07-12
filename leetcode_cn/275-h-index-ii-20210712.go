package leetcode_cn

func hIndex(citations []int) int {
	i := 0
	for i < len(citations) {
		if citations[len(citations)-i-1] >= i+1 {
			i++
			continue
		} else {
			break
		}
	}
	return i
}
