package leetcode_cn

func climbStairs(n int) int {
	a, b, c := 1, 0, 1
	for i := 0; i < n; i++ {
		a, b = b, c
		c = a + b
	}
	return c
}
