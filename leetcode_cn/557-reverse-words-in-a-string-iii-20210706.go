package leetcode_cn

import "strings"

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	res := []byte{}
	for _, str := range strs {
		res = append(res, reverse(str)...)
		res = append(res, ' ')
	}
	return string(res[:len(res)-1])
}

func reverse(s string) []byte {
	b := []byte(s)
	var temp byte
	for i := 0; i < len(b)/2; i++ {
		temp = b[i]
		b[i] = b[len(b)-i-1]
		b[len(b)-i-1] = temp
	}
	return b
}
