package leetcode_cn

func titleToNumber(columnTitle string) int {
	sum := 0
	for i := 0; i < len(columnTitle); i++ {
		sum = sum*26 + int(columnTitle[i]-'A'+1)
	}
	return sum
}
