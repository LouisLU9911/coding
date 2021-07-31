package leetcode_cn

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if canVisit(&grid, &visited, i, j) {
				area := 0
				visit(&grid, &visited, i, j, &area)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func visit(grid *[][]int, visited *[][]bool, i, j int, maxValue *int) {
	if canVisit(grid, visited, i, j) {
		*maxValue += 1
		(*visited)[i][j] = true
		visit(grid, visited, i-1, j, maxValue)
		visit(grid, visited, i+1, j, maxValue)
		visit(grid, visited, i, j-1, maxValue)
		visit(grid, visited, i, j+1, maxValue)
	}
}

func canVisit(grid *[][]int, visited *[][]bool, i, j int) bool {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) || (*grid)[i][j] == 0 || (*visited)[i][j] {
		return false
	}
	return true
}
