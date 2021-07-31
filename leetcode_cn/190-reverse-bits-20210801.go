package leetcode_cn

func reverseBits(num uint32) uint32 {
	var ret uint32 = 0
	var bit uint32 = 1
	for i := 0; i < 32; i++ {
		thisBit := bit & (num >> i)
		ret |= thisBit << (31 - i)
	}
	return ret
}
