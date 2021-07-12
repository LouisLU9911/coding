package leetcode_cn

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
		return nil
	}
	if image[sr][sc] == newColor {
		return image
	}
	doFloodFill(&image, sr, sc, image[sr][sc], newColor)
	return image
}

func doFloodFill(image *[][]int, sr, sc, origin, newColor int) {
	(*image)[sr][sc] = newColor
	// left
	if sc-1 >= 0 && (*image)[sr][sc-1] == origin {
		doFloodFill(image, sr, sc-1, origin, newColor)
	}
	// up
	if sr-1 >= 0 && (*image)[sr-1][sc] == origin {
		doFloodFill(image, sr-1, sc, origin, newColor)
	}
	// right
	if sc+1 < len((*image)[0]) && (*image)[sr][sc+1] == origin {
		doFloodFill(image, sr, sc+1, origin, newColor)
	}
	// down
	if sr+1 < len(*image) && (*image)[sr+1][sc] == origin {
		doFloodFill(image, sr+1, sc, origin, newColor)
	}
}
