package leetcode_cn

func reverseString(s []byte) {
	var temp byte
	for i := 0; i < len(s)/2; i++ {
		temp = s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp
	}
}
