package leetcode_cn

func networkDelayTime(times [][]int, n int, k int) int {
	timeMap := make(map[int]map[int]int)
	for i := 0; i < len(times); i++ {
		u, v, w := times[i][0], times[i][1], times[i][2]
		if _, ok := timeMap[u]; !ok {
			timeMap[u] = make(map[int]int)
		}
		timeMap[u][v] = w
	}
	total := 0
	visited := make([]bool, n+1)
	minTime := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		minTime[i] = 101
	}

	queue := []int{k}
	minTime[k] = 0
	visited[k] = true
	total += 1
	var p int
	for {
		if len(queue) == 0 {
			break
		}
		p = queue[0]
		if v2w, ok := timeMap[p]; ok {
			thisTime := minTime[p]
			for v, w := range v2w {
				if !visited[v] {
					total += 1
					visited[v] = true
				}
				if thisTime+w < minTime[v] {
					queue = append(queue, v)
					minTime[v] = thisTime + w
				}
			}
		}
		queue = queue[1:]
	}

	if total < n {
		return -1
	}

	maxTime := 0
	for i := 1; i < n+1; i++ {
		if minTime[i] > maxTime {
			maxTime = minTime[i]
		}
	}
	return maxTime
}
