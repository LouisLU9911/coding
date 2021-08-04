package leetcode_cn

func eventualSafeNodes(graph [][]int) []int {
	color := make([]int, len(graph))
	for i, _ := range graph {
		// continue if node is loop node or is safe node
		if color[i] == 1 || color[i] == 2 {
			continue
		}
		color[i] = 1
		gray := []int{i}
		dfs(&graph, &gray, &color)
	}
	ret := []int{}
	for i, v := range color {
		if v == 2 {
			ret = append(ret, i)
		}
	}
	return ret
}

func dfs(graph *[][]int, gray *[]int, color *[]int) bool {
	thisNode := (*gray)[len(*gray)-1]
	for _, v := range (*graph)[thisNode] {
		if (*color)[v] == 0 {
			// visit this node
			(*color)[v] = 1
			*gray = append(*gray, v)
			hasLoop := dfs(graph, gray, color)
			if hasLoop {
				return true
			}
			*gray = (*gray)[:len(*gray)-1]
		} else if (*color)[v] == 1 {
			// has loop
			return true
		}
	}
	// is safe node
	(*color)[thisNode] = 2
	return false
}
