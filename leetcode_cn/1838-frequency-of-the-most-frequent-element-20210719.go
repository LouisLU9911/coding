package leetcode_cn

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	l, r, maxFre, remain := 0, 1, 1, k
	for r < len(nums) {
		if r <= l {
			remain = k
			r++
		} else if (nums[r]-nums[r-1])*(r-l) <= remain {
			remain -= ((nums[r] - nums[r-1]) * (r - l))
			if newFre := r - l + 1; newFre > maxFre {
				maxFre = newFre
			}
			r++
		} else {
			remain += (nums[r-1] - nums[l])
			l++
		}
	}
	return maxFre
}
