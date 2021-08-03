package leetcode_cn

func findUnsortedSubarray(nums []int) int {
	minIndex, maxIndex := -1, -1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if minIndex == -1 || nums[i] < nums[minIndex] {
				minIndex = i
			}
			if maxIndex == -1 || nums[i-1] > nums[maxIndex] {
				maxIndex = i - 1
			}
		}
	}
	if minIndex == -1 && maxIndex == -1 {
		return 0
	}
	left := 0
	for left < len(nums) && nums[left] <= nums[minIndex] {
		left++
	}
	right := len(nums) - 1
	for right >= 0 && nums[right] >= nums[maxIndex] {
		right--
	}
	ret := right - left + 1
	return ret
}
