package leetcode_cn

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r && l >= 0 && r < len(nums) {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// if nums[mid] is the first element larger than the target
			if (mid > 0 && nums[mid-1] < target) || mid == 0 {
				return mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// if no position in this list fit the condition, insert into the end
	return len(nums)
}
