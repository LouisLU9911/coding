package leetcode_cn

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	adjacentMap := make(map[int][]int, n)
	for i := 0; i < len(adjacentPairs); i++ {
		adjacentMap[adjacentPairs[i][0]] = append(adjacentMap[adjacentPairs[i][0]], adjacentPairs[i][1])
		adjacentMap[adjacentPairs[i][1]] = append(adjacentMap[adjacentPairs[i][1]], adjacentPairs[i][0])
	}
	begin := 0
	for k, v := range adjacentMap {
		if len(v) == 1 {
			begin = k
			break
		}
	}
	num := []int{begin, adjacentMap[begin][0]}
	for i := adjacentMap[begin][0]; len(adjacentMap[i]) != 1; {
		var next int
		if adjacentMap[i][0] == num[len(num)-2] {
			next = adjacentMap[i][1]
		} else {
			next = adjacentMap[i][0]
		}
		num = append(num, next)
		i = next
	}
	return num
}
