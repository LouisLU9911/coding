package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if i2, ok := numMap[v]; ok {
			return []int{i, i2}
		} else {
			numMap[target-v] = i
		}
	}
	return []int{}
}
