package leetcode_cn

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < (len(nums) >> 1); i++ {
		if sum := nums[i] + nums[len(nums)-i-1]; sum > ans {
			ans = sum
		}
	}
	return ans
}
