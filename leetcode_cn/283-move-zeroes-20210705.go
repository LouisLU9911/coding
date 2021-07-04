package leetcode_cn

func moveZeroes(nums []int) {
	now := 0
	for pos := 0; pos < len(nums); pos++ {
		if nums[pos] != 0 {
			nums[now] = nums[pos]
			now++
		}
	}
	for now < len(nums) {
		nums[now] = 0
		now++
	}
}
