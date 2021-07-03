package leetcode_cn

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r && r < len(nums) && l >= 0 {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
