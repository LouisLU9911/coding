package jianzhi_offer

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum, maxNum := 0, nums[0]
	for pos := 0; pos < len(nums); pos++ {
		sum += nums[pos]
		if sum > maxNum {
			maxNum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxNum
}
