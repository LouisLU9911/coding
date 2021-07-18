package lcp

func numWays(n int, relation [][]int, k int) int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}

	for i := 0; i < len(relation); i++ {
		mat[relation[i][0]][relation[i][1]] = 1
	}

	row := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = mat[0][i]
	}

	res := make([]int, n)
	for i := 0; i < k-1; i++ {

		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += row[k] * mat[k][j]
			}
			res[j] = sum
		}

		for j := 0; j < n; j++ {
			row[j] = res[j]
		}

	}
	return row[n-1]
}
