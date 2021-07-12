package leetcode_cn

import "sort"

func hIndex(citations []int) (h int) {
	sort.Ints(citations)
	i := 0
	for i < len(citations) {
		if citations[len(citations)-1-i] >= i+1 {
			i++
			continue
		} else {
			break
		}
	}
	return i
}
