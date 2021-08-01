package leetcode_cn

func kWeakestRows(mat [][]int, k int) []int {
	val2idx := make([][]int, len(mat[0])+1)
	for i := 0; i < len(mat); i++ {
		sum := 0
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != 0 {
				sum += mat[i][j]
			} else {
				break
			}
		}
		val2idx[sum] = append(val2idx[sum], i)
	}

	ret := []int{}
	for i := 0; i < len(mat[0])+1; i++ {
		for _, v := range val2idx[i] {
			ret = append(ret, v)
			if len(ret) == k {
				return ret
			}
		}
	}
	return ret
}
