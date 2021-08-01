package leetcode_cn

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	ret := make([][]int, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ret[i][j] = 10000
			if matrix[i][j] == 0 {
				ret[i][j] = 0
				continue
			}
			if i > 0 {
				ret[i][j] = min(ret[i][j], ret[i-1][j]+1)
			}
			if j > 0 {
				ret[i][j] = min(ret[i][j], ret[i][j-1]+1)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			if i < m-1 {
				ret[i][j] = min(ret[i][j], ret[i+1][j]+1)
			}
			if j < n-1 {
				ret[i][j] = min(ret[i][j], ret[i][j+1]+1)
			}
		}
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
