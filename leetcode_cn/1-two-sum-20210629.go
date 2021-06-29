package leetcode_cn

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := range nums {
		if v, ok := hashMap[target-nums[i]]; ok {
			return []int{v, i}
		}
		hashMap[nums[i]] = i
	}
	return nil
}
