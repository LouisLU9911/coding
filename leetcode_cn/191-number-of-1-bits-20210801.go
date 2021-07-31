package leetcode_cn

func hammingWeight(num uint32) int {
	cnt := 0
	var bit uint32 = 1
	for i := 0; i < 32; i++ {
		cnt += int(bit & num)
		num = num >> 1
	}
	return cnt
}
