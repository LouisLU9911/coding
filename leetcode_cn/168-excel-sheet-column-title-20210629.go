package leetcode_cn

func convertToTitle(columnNumber int) string {
	quotient := (columnNumber - 1) / 26
	remainder := (columnNumber - 1) % 26
	var res string
	if quotient > 0 {
		prev := convertToTitle(quotient)
		res = prev + string('A'+remainder)
	} else {
		res = string('A' + remainder)
	}
	return res
}
