package leetcode_cn

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for pos := len(nums) - 1; pos >= 0; pos-- {
		lSquare, rSquare := nums[l]*nums[l], nums[r]*nums[r]
		if lSquare > rSquare {
			ret[pos] = lSquare
			l++
		} else {
			ret[pos] = rSquare
			r--
		}
	}
	return ret
}
