package leetcode_cn

func findErrorNums(nums []int) []int {
	sum, rep, flag := 0, 0, true
	for i := 0; i < len(nums) && flag; i++ {
		for nums[i]-1 != i {
			if nums[nums[i]-1]-1 == nums[i]-1 {
				flag = false
				rep = nums[i]
				break
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}

	for _, v := range nums {
		sum += v
	}
	return []int{rep, ((1 + len(nums)) * len(nums) / 2) - sum + rep}
}
